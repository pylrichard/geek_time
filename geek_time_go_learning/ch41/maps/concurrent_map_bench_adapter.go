package maps

import "github.com/easierway/concurrent_map"

type ConcurrentMapBenchAdapter struct {
	cm *concurrent_map.ConcurrentMap
}

func NewConcurrentMapBenchAdapter(numOfPartitions int) *ConcurrentMapBenchAdapter {
	cm := concurrent_map.CreateConcurrentMap(numOfPartitions)

	return &ConcurrentMapBenchAdapter{cm}
}

func (cmba *ConcurrentMapBenchAdapter) Set(key interface{}, value interface{}) {
	cmba.cm.Set(concurrent_map.StrKey(key.(string)), value)
}

func (cmba *ConcurrentMapBenchAdapter) Get(key interface{}) (interface{}, bool) {
	return cmba.cm.Get(concurrent_map.StrKey(key.(string)))
}

func (cmba *ConcurrentMapBenchAdapter) Del(key interface{}) {
	cmba.cm.Del(concurrent_map.StrKey(key.(string)))
}