package pipeline

import "fmt"

// provides understanding of the concept of line
func Idea() {

	add := func(arr []int, value int) []int {
		result := make([]int, len(arr))

		for i, v := range arr {
			result[i] = v + value
		}

		return result
	}

	multiply := func(arr []int, value int) []int {
		result := make([]int, len(arr))

		for i, v := range arr {
			result[i] = v * value
		}

		return result
	}

	list := []int{1, 2, 3}
	fmt.Println(
		multiply(
			add(list, 1),
			2,
		),
	)
}
