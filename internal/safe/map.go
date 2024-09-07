package safe

import (
	"maps"
	"sync"
)

// SyncMap is a thread-safe map
type SyncMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

// NewSyncMap creates a new thread-safe map
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

func (m *SyncMap[K, V]) Store(k K, v V) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *SyncMap[K, V]) Load(k K) (v V, ok bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok = m.m[k]
	return
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m *SyncMap[K, V]) LoadOrStore(k K, v V) (actual V, loaded bool) {
	m.Lock()
	defer m.Unlock()
	actual, loaded = m.m[k]
	if loaded {
		return
	}
	m.m[k] = v
	return v, false
}

// LoadAndDelete deletes the value for a key, and returns the previous value if any.
func (m *SyncMap[K, V]) LoadAndDelete(k K) (v V, loaded bool) {
	m.Lock()
	defer m.Unlock()
	v, loaded = m.m[k]
	if loaded {
		delete(m.m, k)
	}
	return
}

func (m *SyncMap[K, V]) Delete(k K) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *SyncMap[K, V]) Clear() {
	m.Lock()
	defer m.Unlock()
	clear(m.m)
}

func (m *SyncMap[K, V]) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *SyncMap[K, V]) Clone() map[K]V {
	m.RLock()
	defer m.RUnlock()
	return maps.Clone(m.m)
}

func (m *SyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
}
