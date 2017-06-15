package main

// go test -bench=. -benchmem

import "testing"

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}
func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}
