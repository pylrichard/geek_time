package maps

import (
	"strconv"
	"sync"
	"testing"
)

const (
	NumOfReader = 100
	NumOfWriter = 200
)

type Map interface {
	Set(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchMap(b *testing.B, m Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < NumOfWriter; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					m.Set(strconv.Itoa(i), i * i)
					m.Set(strconv.Itoa(i), i * i)
					m.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < NumOfReader; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					m.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	b.Run("RWLockMap", func(b *testing.B) {
		m := NewRWLockMap()
		benchMap(b, m)
	})
	b.Run("SyncMap", func(b *testing.B) {
		m := NewSyncMapBenchAdapter()
		benchMap(b, m)
	})
	b.Run("ConcurrentMap", func(b *testing.B) {
		m := NewConcurrentMapBenchAdapter(199)
		benchMap(b, m)
	})
}