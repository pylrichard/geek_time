package lock

import (
	"log"
	"sync"
)

var cache map[string]string

const (
	NumOfReader = 40
	ReadTimes = 100000
)

func init() {
	cache = make(map[string]string)

	cache["a"] = "aa"
	cache["b"] = "bb"
}

func lockFreeAccess() {
	var wg sync.WaitGroup
	wg.Add(NumOfReader)
	for i := 0; i < NumOfReader; i++ {
		go func() {
			for j := 0; j < ReadTimes; j++ {
				_, err := cache["a"]
				if !err {
					log.Println("Nothing")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func lockAccess() {
	var wg sync.WaitGroup
	wg.Add(NumOfReader)
	m := new(sync.RWMutex)
	for i := 0; i < NumOfReader; i++ {
		go func() {
			for j := 0; j < ReadTimes; j++ {
				m.RLock()
				_, err := cache["a"]
				if !err {
					log.Println("Nothing")
				}
				m.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}