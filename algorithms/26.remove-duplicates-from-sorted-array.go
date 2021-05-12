package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	behindIdx := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if nums[i] == nums[behindIdx] {
			continue
		}
		nums[behindIdx+1] = nums[i]
		behindIdx = behindIdx + 1
	}
	nums = nums[:behindIdx+1]
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
