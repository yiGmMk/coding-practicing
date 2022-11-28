package main

//
// void foo() {
// }
//
import "C"
import "testing"

func CallCFunc() {
	C.foo()
}

func foo() {
}

func CallGoFunc() {
	foo()
}

// chapter9/sources/go-cgo/cgo-perf/cgo_call_test.go

func BenchmarkCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallCFunc()
	}
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallGoFunc()
	}
}
