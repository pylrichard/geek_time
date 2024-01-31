package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type counter struct {
	num     uint
	rwMutex sync.RWMutex
}

// getNum 返回当前计数
func (c *counter) getNum() uint {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	return c.num
}

// add 增加计数器的值，并会返回增加后的计数
func (c *counter) add(increment uint) uint {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.num += increment

	return c.num
}

func main() {
	c := counter{}
	count(&c)
	redundantUnlock()
}

func count(c *counter) {
	sign := make(chan struct{}, 3)
	// 增加计数
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number in counter: %d [%d-%d]",
				c.getNum(), 1, j)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= 20; k++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("The number in counter: %d [%d-%d]",
				c.getNum(), 2, k)
		}
	}()
	<-sign
	<-sign
	<-sign
}

func redundantUnlock() {
	var rwMutex sync.RWMutex

	// 示例1
	// 这里会引发panic
	// rwMutex.Unlock()

	// 示例2
	// 这里会引发panic
	// rwMutex.RUnlock()

	// 示例3
	rwMutex.RLock()
	fmt.Println("hello world")
	// 这里会引发panic
	// rwMutex.Unlock()
	rwMutex.RUnlock()

	// 示例4
	rwMutex.Lock()
	fmt.Println("hello world")
	// 这里会引发panic
	// rwMutex.RUnlock()
	rwMutex.Unlock()
}
