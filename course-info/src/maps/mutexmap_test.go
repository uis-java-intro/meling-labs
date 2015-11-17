package maps

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexMap(t *testing.T) {
	m := &mutexMap{make(map[string]string), &sync.Mutex{}}

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				x := m.lookup("mutexmap")
				fmt.Println(x)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func(j int) {
			val := fmt.Sprintf("mutex %d", j)
			for {
				m.insert("mutexmap", val)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	timeout := time.After(5 * time.Second)
	<-timeout
}
