package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("put %d\n", i)
			//写入数据会阻塞，receiver读取后继续写入
			ch<-i
		}
		fmt.Println("producer close channel")
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(idx int, ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("receiver %d get %d\n", idx, data)
			} else {
				fmt.Printf("receiver %d break loop\n", idx)
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 2)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(1, ch, &wg)
	/*
		增加receiver随机读取数据
	 */
	wg.Add(1)
	dataReceiver(2, ch, &wg)
	wg.Wait()
}