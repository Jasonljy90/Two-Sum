package main

import "fmt"

func sum(nums [4]int, target int) [2]int {
	var sum int
	var indices [2]int
	for i := 0; i < (len(nums) - 1); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == target {
				indices = [2]int{i, j}
				return indices
			}
		}
	}
	return indices
}

func main() {
	nums := [4]int{2, 7, 11, 15}
	s := sum(nums, 17)
	fmt.Println(s)
}
