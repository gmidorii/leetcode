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
	for xi := 0; xi < len(nums); xi++ {
		for yi := xi + 1; yi < len(nums); yi++ {
			zn := -1 * (nums[xi] + nums[yi])
			// if m(nums[xi], nums[yi], zn, result) {
			// 	continue
			// }
			for zi := yi + 1; zi < len(nums); zi++ {
				if zn == nums[zi] {
					result = append(result, []int{nums[xi], nums[yi], nums[zi]})
					break
				}
			}
		}
	}

	// TOCO: result の重複除去のほうが検索回数が少なくなるのでは
	var res [][]int
	for _, rr := range result {
		if !match(rr[0], rr[1], rr[2], res) {
			res = append(res, rr)
		}
	}
	return res
}

func match(x, y, z int, res [][]int) bool {
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
