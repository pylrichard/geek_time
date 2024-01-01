package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var m sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				m.Unlock()
			}()
			m.Lock()
			counter++
		}()
	}
	//避免子协程未执行完毕，主协程退出
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var m sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				m.Unlock()
			}()
			m.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}