package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	fmt.Println("service start")
	time.Sleep(50 * time.Millisecond)

	return "service done"
}

func otherService() {
	fmt.Println("otherService start")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("otherService done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherService()
}

func AsyncService() chan string {
	ch := make(chan string)
	//缓冲Channel容量为1，放入元素后可以立即返回，实现异步
	//ch := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("service return result")
		//不使用缓冲Channel，会在此处阻塞等待读取
		ch<-ret
		fmt.Println("asyncService goroutine done")
	}()

	return ch
}

func TestAsyncService(t *testing.T) {
	ch := AsyncService()
	otherService()
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}