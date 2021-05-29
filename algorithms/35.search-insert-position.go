package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, n := range nums {
		if n >= target {
			return i
		}
	}
	return len(nums)
}

func searchInsert2(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}
	left := 0
	right := len(nums) - 1
	for {
		mid := left + (right-left)/2
		midP := mid + 1
		fmt.Printf("l:%v, m:%v, mp:%v, r:%v\n", left, mid, midP, right)
		fmt.Printf("v l:%v, m:%v, mp:%v, r:%v\n", nums[left], nums[mid], nums[midP], nums[right])
		if target > nums[mid] && target <= nums[midP] {
			return midP
		}
		if midP == len(nums)-1 {
			return len(nums)
		}
		if mid == 0 {
			return 0
		}
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return len(nums)
}
