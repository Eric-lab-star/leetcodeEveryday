package ans

import (
	"fmt"
)

var test1 = []int{1, 2, 3, 4}

func Submit() {
	result := minimumOperations(test1)
	fmt.Printf("result: %d\n", result)
}

func minimumOperations(nums []int) int {
	result := 0
	for _, v := range nums {
		modulo := v % 3
		if modulo != 0 {
			result++
		}
	}
	return result
}
