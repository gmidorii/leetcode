package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			nums[0] = -1
			return 0
		} else {
			return 1
		}
	}
	behindIdx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[behindIdx+1] = nums[i]
		behindIdx++
	}
	return behindIdx + 1
}
