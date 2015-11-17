package spinlock

import "testing"

var spinner SpinLock
var ospinner SpinLock
var value []int32
var ovalue []int32

func inc(val int32) {
	spinner.Lock()
	val++
	spinner.Unlock()
}

func oinc(val int32) {
	ospinner.Lock()
	val++
	ospinner.Unlock()
}

func BenchmarkSpin1(b *testing.B) {
	value = make([]int32, b.N)
	done := make(chan bool)
	b.Log(b.N)
	go func() {
		for i := 0; i < b.N; i++ {
			inc(value[i])
		}
		done <- true
	}()
	<-done
}

func BenchmarkSpin2(b *testing.B) {
	value = make([]int32, b.N)
	done := make(chan bool)
	go func() {
		for i := 0; i < b.N; i++ {
			inc(value[i])
		}
		done <- true
	}()

	ovalue = make([]int32, b.N)
	odone := make(chan bool)
	go func() {
		for i := 0; i < b.N; i++ {
			inc(ovalue[i])
		}
		odone <- true
	}()
	<-done
	<-odone
}
