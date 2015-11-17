package maps

import (
	"fmt"
	"testing"
	"time"
)

func TestRWLockMap(t *testing.T) {
	m := &rwLockMap{make(map[string]string), NewRWLock()}

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				x := m.lookup("rwlock")
				fmt.Println(x)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func(j int) {
			val := fmt.Sprintf("rwlock %d", j)
			for {
				m.insert("rwlock", val)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	timeout := time.After(5 * time.Second)
	<-timeout
}
