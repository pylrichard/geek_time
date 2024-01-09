package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type DemoCollector struct {
	receiver	EventReceiver
	ctx			context.Context
	stopChan	chan struct{}
	name		string
	content		string
}

func NewDemoCollector(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan:	make(chan struct{}),
		name:		name,
		content:	content,
	}
}

func (dc *DemoCollector) Init(receiver EventReceiver) error {
	fmt.Println("init collector ", dc.name)
	dc.receiver = receiver

	return nil
}

func (dc *DemoCollector) Start(ctx context.Context) error {
	fmt.Println("start collector ", dc.name)
	for {
		select {
		case <-ctx.Done():
			dc.stopChan <-struct{}{}
			break
		default:
			time.Sleep(50 * time.Millisecond)
			dc.receiver.OnEvent(Event{dc.name, dc.content})
		}
	}
}

func (dc *DemoCollector) Stop() error {
	fmt.Println("stop collector ", dc.name)
	select {
	case <-dc.stopChan:
		return nil
	case <-time.After(1 * time.Second):
		return errors.New("failed to stop for timeout")
	}
}

func (dc *DemoCollector) Destroy() error {
	fmt.Println(dc.name, "released resources")

	return nil
}