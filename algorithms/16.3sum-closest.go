package main

func threeSumClosest(nums []int, target int) int {
	closest := (10 * 10 * 10) * 3
	for xi := 0; xi < len(nums); xi++ {
		for yi := xi + 1; yi < len(nums); yi++ {
			for zi := yi + 1; zi < len(nums); zi++ {
				c := target - (nums[xi] + nums[yi] + nums[zi])
				if c < 0 {
					c = -1 * c
				}
				d := target - closest
				if d < 0 {
					d = -1 * d
				}
				if c < d {
					closest = (nums[xi] + nums[yi] + nums[zi])
				}
			}
		}
	}

	return closest
}
