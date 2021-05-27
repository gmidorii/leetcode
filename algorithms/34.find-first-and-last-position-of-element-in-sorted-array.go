package main

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			l := mid
			r := mid
			for {
				if (l < 0 || nums[l] != target) && (r >= len(nums) || nums[r] != target) {
					return []int{l + 1, r - 1}
				}
				if l >= 0 && nums[l] == target {
					l--
				}
				if r < len(nums) && nums[r] == target {
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
