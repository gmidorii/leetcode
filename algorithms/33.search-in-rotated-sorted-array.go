package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}
