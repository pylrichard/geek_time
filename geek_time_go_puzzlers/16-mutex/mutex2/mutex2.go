package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// singleHandler 代表单次处理函数的类型
type singleHandler func() (data string, n int, err error)

// handlerConfig 代表处理流程配置的类型
type handlerConfig struct {
	// 单次处理函数
	handler singleHandler
	// 需要启用的goroutine数量
	goNum int
	// 单个goroutine中的处理次数
	number int
	// 单个goroutine中的处理间隔时间
	interval time.Duration
	// 数据量计数器，以字节为单位
	counter int
	// 数据量计数器专用的互斥锁
	counterMutex sync.Mutex
}

// count 会增加计数器的值，并会返回增加后的计数
func (hc *handlerConfig) count(increment int) int {
	hc.counterMutex.Lock()
	defer hc.counterMutex.Unlock()
	hc.counter += increment
	return hc.counter
}

func main() {
	// 在下面函数中直接使用，不要传递
	var mu sync.Mutex

	// genWriter代表用于生成写入函数的函数
	genWriter := func(writer io.Writer) singleHandler {
		return func() (data string, n int, err error) {
			// 准备数据
			data = fmt.Sprintf("%s\t",
				time.Now().Format(time.StampNano))
			// 写入数据
			mu.Lock()
			defer mu.Unlock()
			n, err = writer.Write([]byte(data))
			return
		}
	}

	// genReader代表用于生成读取函数的函数
	genReader := func(reader io.Reader) singleHandler {
		return func() (data string, n int, err error) {
			buffer, ok := reader.(*bytes.Buffer)
			if !ok {
				err = errors.New("unsupported reader")
				return
			}
			// 读取数据
			mu.Lock()
			defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}

	var buffer bytes.Buffer

	// 数据写入配置
	writeConfig := handlerConfig{
		handler:  genWriter(&buffer),
		goNum:    5,
		number:   4,
		interval: time.Millisecond * 100,
	}
	// 数据读取配置
	readConfig := handlerConfig{
		handler:  genReader(&buffer),
		goNum:    10,
		number:   2,
		interval: time.Millisecond * 100,
	}

	sign := make(chan struct{}, writeConfig.goNum+readConfig.goNum)

	// 启用多个goroutine对缓冲区进行多次数据写入
	for i := 1; i <= writeConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= writeConfig.number; j++ {
				time.Sleep(writeConfig.interval)
				data, n, err := writeConfig.handler()
				if err != nil {
					log.Printf("writer [%d-%d]: error: %s",
						i, j, err)
					continue
				}
				total := writeConfig.count(n)
				log.Printf("writer [%d-%d]: %s (total: %d)",
					i, j, data, total)
			}
		}(i)
	}

	// 启用多个goroutine对缓冲区进行多次数据读取
	for i := 1; i <= readConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= readConfig.number; j++ {
				time.Sleep(readConfig.interval)
				var data string
				var n int
				var err error
				for {
					data, n, err = readConfig.handler()
					if err == nil || err != io.EOF {
						break
					}
					// 如果读比写快（读时会发生EOF错误），那就等一会儿再读。
					time.Sleep(readConfig.interval)
				}
				if err != nil {
					log.Printf("reader [%d-%d]: error: %s",
						i, j, err)
					continue
				}
				total := readConfig.count(n)
				log.Printf("reader [%d-%d]: %s (total: %d)",
					i, j, data, total)
			}
		}(i)
	}

	// signNumber代表需要接收的信号数量
	signNumber := writeConfig.goNum + readConfig.goNum
	// 等待上面启用的所有goroutine的运行全部结束
	for j := 0; j < signNumber; j++ {
		<-sign
	}
}
