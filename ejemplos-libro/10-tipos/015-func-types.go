package main

import "math/rand"

type Generator func() int

func GenerateAll(gens ...func() int) []int {
	nums := make([]int, 0, len(gens))
	for _, g := range gens {
		nums = append(nums, g())
	}
	return nums
}

func ZeroGenerator() int {
	return 0
}

func CounterGenerator() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func RandomGenerator(seed int64) func() int {
	rnd := rand.NewSource(seed)
	return func() int {
		return int(rnd.Int63())
	}
}

func main() {

}
