package utils

import (
	"sync"
	"time"
)

type ExpiringMap struct {
	mu  sync.RWMutex
	m   map[string]string
	ttl time.Duration
}

func NewExpiringMap(ttl time.Duration) *ExpiringMap {
	em := &ExpiringMap{
		m:   make(map[string]string),
		ttl: ttl,
	}
	return em
}

func (em *ExpiringMap) Del(key, value string) {
	em.mu.Lock()
	defer em.mu.Unlock()
	if v, ok := em.m[key]; ok && v == value {
		delete(em.m, key)
	}
}

func (em *ExpiringMap) Set(key, value string) {
	em.mu.Lock()
	em.m[key] = value
	em.mu.Unlock()

	time.AfterFunc(em.ttl, func() {
		em.Del(key, value)
	})
}

func (em *ExpiringMap) Get(key string) (value string, ok bool) {
	em.mu.RLock()
	defer em.mu.RUnlock()
	value, ok = em.m[key]
	return
}
