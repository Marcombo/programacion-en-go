package funciones

import "testing"

func TestFib_0(t *testing.T) {
	res := Fib(0)
	if len(res) != 1 || res[0] != 1 {
		t.Errorf("fib(0) expected to be [1]. Got %+v", res)
	}
}

func TestFib_1(t *testing.T) {
	res := Fib(1)
	if len(res) != 2 || res[0] != 1 || res[1] != 1 {
		t.Errorf("fib(1) expected to be [1 1]. Got %+v", res)
	}
}

func TestFib_N(t *testing.T) {
	res := Fib(7)
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21}
	if len(res) != len(expected) {
		t.Errorf("fib(7) expected to be %+v. Got %+v", expected, res)
		return
	}
	for i := range expected {
		if expected[i] != res[i] {
			t.Errorf("fib(7) expected to be %+v. Got %+v", expected, res)
			return
		}
	}
}

/* ejecutar con:
go test -bench=.
go test -bench=. -benchmem
go test -bench=. -benchmem -benchtime=10s
*/
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}
