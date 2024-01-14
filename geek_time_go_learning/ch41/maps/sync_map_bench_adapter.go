package maps

import "sync"

type SyncMapBenchAdapter struct {
	m sync.Map
}

func NewSyncMapBenchAdapter() *SyncMapBenchAdapter {
	return &SyncMapBenchAdapter{}
}

func (smba *SyncMapBenchAdapter) Get(key interface{}) (interface{}, bool) {
	return smba.m.Load(key)
}

func (smba *SyncMapBenchAdapter) Set(key interface{}, value interface{}) {
	smba.m.Store(key, value)
}

func (smba *SyncMapBenchAdapter) Del(key interface{}) {
	smba.m.Delete(key)
}