package twosum

func TwoSum(nums []int, target int) [2]int {
	for i, num := range nums {
		for j, n := range nums {
			if i != j && n+num == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}
