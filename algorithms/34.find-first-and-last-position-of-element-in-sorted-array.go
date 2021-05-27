package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("%v %v %v\n", nums[left], nums[mid], nums[right])
		if nums[mid] == target {
			l := mid
			r := mid
			for {
				if nums[l] != target && nums[r] != target {
					return []int{l + 1, r - 1}
				}
				if nums[l] == target {
					l--
				}
				if nums[r] == target {
					r++
				}
			}
		}
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
