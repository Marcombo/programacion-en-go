package functions

import "testing"

/* run with:
go test -bench=.
go test -bench=. -benchmem
go test -bench=. -benchmem -benchtime=10s
*/
func BenchmarkFib(b *testing.B) {
	// use b.StartTimer() to ignore time-expensive initialization functions
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}
