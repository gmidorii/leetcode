package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	lastIdx := len(nums) - 1
	for i := lastIdx; i > 0; i-- {
		if nums[i] > nums[i-1] {
			max := 101
			maxIdx := -1
			for j := (i - 1) + 1; j < len(nums); j++ {
				if nums[j] <= nums[i-1] {
					continue
				}
				if nums[j] < max {
					max = nums[j]
					maxIdx = j
				}
			}
			nums[i-1], nums[maxIdx] = swap(nums[i-1], nums[maxIdx])
			sortSkipIdx(nums, i)
			return
		}
	}
	reverse(nums)
}

func swap(a, b int) (int, int) {
	return b, a
}

func sortSkipIdx(nums []int, firstIdx int) {
	if len(nums) <= 1 {
		return
	}
	lastIdx := len(nums) - 1
	for {
		if lastIdx < firstIdx {
			break
		}
		for i := firstIdx; i < lastIdx; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = swap(nums[i], nums[i+1])
			}
		}
		lastIdx--
	}
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = swap(nums[i], nums[len(nums)-i-1])
	}
}
