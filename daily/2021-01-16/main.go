package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{3, 2, 1, 5, 6, 4}
	k := 2
	expected := 5
	output := findKthLargest(input, k)

	fmt.Println(output == expected)
}

func findKthLargest(nums []int, k int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}
