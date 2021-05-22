package main

/*
1. 末尾から順番に比較して若いほうが小さい値の場合は入れ替えて返却
2. 先頭との比較で先頭のほうが小さい場合は、先頭の値より最も近く大きい値と入れ替えて、小さい順にソート
3. 先頭との比較で先頭のほうが大きい場合は、逆順にして返却
*/

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	first := nums[0]
	max := 101
	maxIdx := len(nums) - 1
	maxNext := 101
	maxNextIdx := len(nums) - 1
	lastIdx := len(nums) - 1
	for i := lastIdx; i > 0; i-- {
		if nums[i] >= first {
			if nums[i] < max {
				maxNext = max
				maxNextIdx = maxIdx
				max = nums[i]
				maxIdx = i
			} else if nums[i] < maxNext {
				maxNext = nums[i]
				maxNextIdx = i
			}
		}
		if nums[i] > nums[i-1] {
			if i == 1 && nums[i] == max {
				nums[i-1], nums[maxNextIdx] = swap(nums[i-1], nums[maxNextIdx])
				sortSkipIdx(nums, i)
				return
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
