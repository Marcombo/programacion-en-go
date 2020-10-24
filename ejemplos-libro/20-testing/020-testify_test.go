package functions

/**
To import testify, use:
	go get -t github.com/stretchr/testify/assert
	govendor
	go mod (recommended(
*/
import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestFib_0_Testify(t *testing.T) {
	assert.Equal(t, []int{1}, Fib(0))
}

func TestFib_1_Testify(t *testing.T) {
	assert.Equal(t, []int{1, 1}, Fib(1))
}

func TestFib_N_Testify(t *testing.T) {
	assert.Equal(t, []int{1, 1, 2, 3, 5, 8, 13, 21}, Fib(7))
}

// Using test cases
func TestFib(t *testing.T) {
	testCases := []struct {
		param    int
		expected []int
	}{
		{0, []int{1}},
		{1, []int{1, 1}},
		{7, []int{1, 1, 2, 3, 5, 8, 13, 21}},
		{10, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}},
	}

	for _, c := range testCases {
		t.Run(fmt.Sprintf("test case fib(%d)", c.param), func(t *testing.T) {
			assert.Equal(t, c.expected, Fib(c.param))
		})
	}
}
