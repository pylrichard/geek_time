package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(500 * time.Millisecond)

	return "service done"
}

func AsyncService() chan string {
	ch := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("service return result")
		ch<-ret
		fmt.Println("asyncService goroutine done")
	}()

	return ch
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Logf("result:%s", ret)
	case <-time.After(100 * time.Millisecond):
		t.Error("time out")
	}
}