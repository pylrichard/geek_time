package cancel_by_close_test

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel1(cancelChan chan struct{}) {
	//通知一个协程
	cancelChan<-struct{}{}
}

func cancel2(cancelChan chan struct{}) {
	//广播实现解耦
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//cancel1(cancelChan)
	cancel2(cancelChan)
	time.Sleep(1 * time.Second)
}