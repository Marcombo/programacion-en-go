package funciones

func Fib(n int) []int {
	if n == 0 {
		return []int{1}
	}
	seq := make([]int, n+1)
	seq[0] = 1
	seq[1] = 1
	for i := 2; i <= n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq
}
