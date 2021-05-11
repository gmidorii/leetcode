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
