package maps

import "sync"

type mutexMap struct {
	ma map[string]string
	mu *sync.Mutex
}

func (m *mutexMap) lookup(key string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.ma[key]
}

func (m *mutexMap) insert(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ma[key] = value
}
