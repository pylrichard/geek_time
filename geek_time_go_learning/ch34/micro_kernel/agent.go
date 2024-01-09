package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("can not take the operation in the current state")

type CollectorsErrors struct {
	Errors []error
}

func (ce CollectorsErrors) Error() string {
	var strArray []string
	for _, err := range ce.Errors {
		strArray = append(strArray, err.Error())
	}

	return strings.Join(strArray, ";")
}

type Event struct {
	Source	string
	Content	string
}

type EventReceiver interface {
	OnEvent(event Event)
}

//Collector 相当于Plugin
type Collector interface {
	//Init 外部创建Collector，在Agent.RegisterCollector()中初始化，事件通知相应实现EventReceiver接口的对象，如Agent
	Init(receiver EventReceiver) error
	//Start Collector运行在不同的协程，在Agent.Stop()调用context.CancelFunc
	Start(ctx context.Context) error
	Stop() error
	Destroy() error
}

//Agent 相当于Kernel
type Agent struct {
	collectors map[string]Collector
	//在协程中处理事件
	eventBuf   chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	//state 通过判断状态避免两次启动，重复执行初始化逻辑
	state      int
}

func NewAgent(size int) *Agent {
	agent := Agent{
		collectors: map[string]Collector{},
		eventBuf:   make(chan Event, size),
		state:      Waiting,
	}

	return &agent
}

func (a *Agent) RegisterCollector(name string, collector Collector) error {
	if a.state != Waiting {
		return WrongStateError
	}
	a.collectors[name] = collector

	return collector.Init(a)
}

func (a *Agent) Start() error {
	if a.state != Waiting {
		return WrongStateError
	}
	a.state = Running
	a.ctx, a.cancel = context.WithCancel(context.Background())
	go a.EventProcessor()

	return a.startCollectors()
}

func (a *Agent) Stop() error {
	if a.state != Running {
		return WrongStateError
	}
	a.state = Waiting
	a.cancel()

	return a.stopCollectors()
}

func (a *Agent) Destroy() error {
	if a.state != Waiting {
		return WrongStateError
	}

	return a.destroyCollectors()
}

func (a *Agent) EventProcessor() {
	var events [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			//OnEvent()中将事件写入eventBuf
			case events[i] = <-a.eventBuf:
			case <-a.ctx.Done():
				return
			}
		}
		fmt.Println(events)
	}
}

func (a *Agent) OnEvent(event Event) {
	//EventProcessor()中读取事件
	a.eventBuf <-event
}

func (a *Agent) startCollectors() error {
	var err error
	var errs CollectorsErrors
	var m sync.Mutex
	for name, collector := range a.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				m.Unlock()
			}()
			err = collector.Start(ctx)
			m.Lock()
			if err != nil {
				errs.Errors = append(errs.Errors, errors.New(name + ":" + err.Error()))
			}
		}(name, collector, a.ctx)
	}
	if len(errs.Errors) == 0 {
		return nil
	}

	return errs
}

func (a *Agent) stopCollectors() error {
	var err error
	var errs CollectorsErrors
	for name, collector := range a.collectors {
		if err = collector.Stop(); err != nil {
			errs.Errors = append(errs.Errors, errors.New(name + ":" + err.Error()))
		}
	}
	if len(errs.Errors) == 0 {
		return nil
	}

	return errs
}

func (a *Agent) destroyCollectors() error {
	var err error
	var errs CollectorsErrors
	for name, collector := range a.collectors {
		if err = collector.Destroy(); err != nil {
			errs.Errors = append(errs.Errors, errors.New(name + ":" + err.Error()))
		}
	}
	if len(errs.Errors) == 0 {
		return nil
	}

	return errs
}