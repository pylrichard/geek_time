package maps

import "sync"

type RWLockMap struct {
	lock	sync.RWMutex
	m		map[interface{}]interface{}
}

func NewRWLockMap() *RWLockMap {
	m := make(map[interface{}]interface{}, 0)

	return &RWLockMap{m: m}
}

func (rwlm *RWLockMap) Get(key interface{}) (interface{}, bool) {
	rwlm.lock.RLock()
	v, ok := rwlm.m[key]
	rwlm.lock.RUnlock()

	return v, ok
}

func (rwlm *RWLockMap) Set(key interface{}, value interface{}) {
	rwlm.lock.Lock()
	rwlm.m[key] = value
	rwlm.lock.Unlock()
}

func (rwlm *RWLockMap) Del(key interface{}) {
	rwlm.lock.Lock()
	delete(rwlm.m, key)
	rwlm.lock.Unlock()
}