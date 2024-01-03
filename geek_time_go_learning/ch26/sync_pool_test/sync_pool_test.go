package sync_pool_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		//P的私有对象和共享池都获取不到对象，调用New()创建新对象
		New: func() interface{} {
			fmt.Println("Create a new obj")

			return 10
		},
	}

	//(int)进行类型断言
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//GC清除sync.Pool缓存的对象
	//runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInRoutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new obj")

			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}