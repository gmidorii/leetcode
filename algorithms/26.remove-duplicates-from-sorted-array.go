package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := len(nums)
	behindIdx := 0
	idx := 1
	for i := 0; i < l; i++ {
		if i == 0 {
			continue
		}
		if idx+1 > len(nums) {
			break
		}
		fmt.Println(nums)
		fmt.Println(behindIdx + 1)
		fmt.Println(len(nums))
		fmt.Println(idx)
		fmt.Println(nums[idx])
		if nums[idx] == nums[behindIdx] {
			idx++
			continue
		}
		fmt.Println("not equal")
		nums = append(nums[:behindIdx+1], nums[idx:]...)
		behindIdx = behindIdx + 1
	}
	nums = nums[:behindIdx+1]
	fmt.Println(nums)
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
