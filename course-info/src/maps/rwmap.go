package maps

import (
	"fmt"
	"sync"
	"time"
)

type syncmap struct {
	ma map[string]string
	mu *sync.Mutex
}

func xTMPCHANGEmain() {
	smap := &syncmap{make(map[string]string), &sync.Mutex{}}

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				x := smap.lookup("hein")
				fmt.Println(x)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		go func(j int) {
			val := fmt.Sprintf("meling %d", j)
			for {
				smap.insert("hein", val)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	ch := make(chan bool)
	<-ch
}

func (m *syncmap) lookup(key string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.ma[key]
}

func (m *syncmap) insert(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ma[key] = value
}
