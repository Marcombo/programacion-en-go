package funciones

import "testing"

type ejemploTest struct {
	num      uint64
	esperado uint64
}

func TestFactorial(t *testing.T) {
	ejemplos := []ejemploTest{
		{num: 1, esperado: 1},
		{num: 7, esperado: 5040},
		{num: 10, esperado: 3628800},
		{num: 13, esperado: 6227020800},
	}
	for _, ej := range ejemplos {
		resultado := Factorial(ej.num)
		if resultado != ej.esperado {
			t.Errorf("Factorial(%d): Esperaba %d. Retornó %d",
				ej.num, ej.esperado, resultado)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Factorial(21)
	}
}
