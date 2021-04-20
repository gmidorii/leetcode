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

	initXi := 0
	initYi := 1
	initZi := 2
	var result [][]int
	m := func(x, y, z int, result [][]int) bool {
		xm := false
		ym := false
		zm := false
		for _, rr := range result {
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
			}
		}
		if xm && ym && zm {
			return true
		}
		return false
	}
	for xi := initXi; xi < len(nums); xi++ {
		for yi := initYi; yi < len(nums); yi++ {
			for zi := initZi; zi < len(nums); zi++ {
				if (nums[xi]+nums[yi]+nums[zi]) == 0 && !m(nums[xi], nums[yi], nums[zi], result) {
					result = append(result, []int{nums[xi], nums[yi], nums[zi]})
				}
			}
			initZi = initZi + 1
		}
		initYi = initZi + 1
	}
	return result
}
