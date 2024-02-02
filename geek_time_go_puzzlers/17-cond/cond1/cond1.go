package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// 0代表信箱空，1代表信箱满
	var mailbox uint8
	// lock代表信箱上的锁
	var lock sync.RWMutex
	// sendCond代表专用于发信的条件变量
	sendCond := sync.NewCond(&lock)
	// recvCond代表专用于收信的条件变量
	recvCond := sync.NewCond(lock.RLocker())
	sign := make(chan struct{}, 3)
	max := 5

	// 用于发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	// 用于收信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
