package task_done_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)

	return fmt.Sprintf("task result from routine %d", id)
}

func GetTaskResult() chan string {
	num := 6
	ch := make(chan string, num)
	//ch := make(chan string)
	for i := 0; i < num; i++ {
		go func(i int) {
			ret := runTask(i)
			//channel中数据没有被读取，写入者会阻塞
			//使用buffered channel可以避免写入者阻塞
			ch<-ret
			fmt.Printf("routine %d exited\n", i)
		}(i)
	}

	return ch
}

func TestAnyTaskDone(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	ch := GetTaskResult()
	t.Log(<-ch)
	//t.Log(<-ch)
	//t.Log(<-ch)

	time.Sleep(1 * time.Second)
	t.Log("After:", runtime.NumGoroutine())
}

func TestAllTaskDone(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	ch := GetTaskResult()
	final := ""
	for i := 0; i < cap(ch); i++ {
		final += <-ch + "\n"
	}
	t.Log(final)
	time.Sleep(1 * time.Second)
	t.Log("After:", runtime.NumGoroutine())
}
