package main

func threeSum(nums []int) [][]int {
	xs := make([]int, len(nums))
	ys := make([]int, len(nums))
	zs := make([]int, len(nums))

	if len(nums) <= 2 {
		return [][]int{}
	}

	copy(xs, nums)
	copy(ys, nums)
	copy(zs, nums)

	var result [][]int
	m := func(x, y, z int, res [][]int) bool {
		xm := false
		ym := false
		zm := false
		for _, rr := range res {
			for _, r := range rr {
				if !xm && r == x {
					xm = true
					continue
				}
				if !ym && r == y {
					ym = true
					continue
				}
				if !zm && r == z {
					zm = true
					continue
				}
				break
			}
			if xm && ym && zm {
				return true
			}
			xm = false
			ym = false
			zm = false
		}

		return false
	}
	mm := func(x, y int, res [][]int) bool {
		xm := false
		ym := false
		for _, rr := range res {
			for _, r := range rr {
				if !xm && r == x {
					xm = true
					continue
				}
				if !ym && r == y {
					ym = true
					continue
				}
				break
			}
			if xm && ym {
				return true
			}
			xm = false
			ym = false
		}
		return false
	}
	for xi := 0; xi < len(nums); xi++ {
		for yi := xi + 1; yi < len(nums); yi++ {
			if mm(nums[xi], nums[yi], result) {
				continue
			}
			for zi := yi + 1; zi < len(nums); zi++ {
				if (nums[xi]+nums[yi]+nums[zi]) == 0 && !m(nums[xi], nums[yi], nums[zi], result) {
					result = append(result, []int{nums[xi], nums[yi], nums[zi]})
					break
				}
			}
		}
	}
	return result
}
