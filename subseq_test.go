// Public domain.

package subseq_test

import (
	"fmt"

	"github.com/soniakeys/subseq"
)

func ExampleLongestIncreasing() {
	a := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	fmt.Println(subseq.LongestIncreasing(a))
	// Output:
	// [0 2 6 9 11 15]
}

func ExampleLongestCommon() {
	a := []int{1, 7, 4, 8, 4, 8, 10, 1, 2, 8, 1}
	b := []int{1, 2, 8, 1, 4, 3, 13, 23, 0, -2, 1, 2, 8, 1, 4, 3, 13}
	fmt.Println(subseq.LongestCommon(a, b))
	// Output:
	// [1 8 4 1 2 8 1]
}
