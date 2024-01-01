package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	//主协程休眠等待子协程执行完毕
	time.Sleep(time.Millisecond * 50)
}

func TestGoroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			//i在子协程中共享
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 50)
}