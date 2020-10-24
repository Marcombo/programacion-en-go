package main

import "fmt"

// Zero arguments, no return function
func Hello() {
	fmt.Println("hello")
}

// Multiple arguments and return type
func CellName(row rune, col int) string {
	return fmt.Sprint(row, col)
}

// Grouping arguments of the same type
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Multiple return values. First returned value: max, second returned value: min
func MaxMin(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

// Named return values. Usually not recomended
func CountUpperLower(text string) (upper int, lower int) {
	upper = 0
	lower = 0
	for _, c := range text {
		if c >= 'A' && c <= 'Z' {
			upper++
		} else if c >= 'a' && c <= 'z' {
			lower++
		}
	}
	return // you can't avoid the return
}

// Varargs function
func AddAll(n ...int) int {
	sum := 0
	for _, s := range n {
		sum += s
	}
	return sum
}

func main() {
	Hello()
	fmt.Println("The top-lef cell is named ", CellName('A', 1))
	fmt.Println("The highest of 5 and 6 is", Max(5, 6))

	ma, mi := MaxMin(4, 7)
	fmt.Printf("%v is higher than %v\n", ma, mi)

	str := "Hello my name is Ralph"
	up, lo := CountUpperLower(str)
	fmt.Printf("%q has %d uppercase and %d lowercase letters\n", str, up, lo)

	fmt.Println("3+4+5+6 =", AddAll(3, 4, 5, 6))

	nums := []int{7, 8, 3}
	all := AddAll(nums...) // spread operator

	fmt.Printf("Adding all numbers in %v: %v\n", nums, all)
}
