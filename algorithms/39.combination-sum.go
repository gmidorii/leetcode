package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var combinations [][]int
	for i := 0; i < len(candidates); i++ {
		c := candidates[i]
		if c > target {
			continue
		}
		var itmp [][]int
		itmp = append(itmp, []int{c})
		for j := i + 1; j < len(candidates); j++ {
			cj := candidates[j]
			if cj > target {
				continue
			}
			jtmp := make([][]int, 0, len(itmp))
			jtmp = append(jtmp, itmp...)
			for _, r := range itmp {
				jtmp = append(jtmp, append(r, cj))
			}
			itmp = jtmp
		}
		combinations = append(combinations, itmp...)
	}

	var res [][]int
	for _, items := range combinations {
		for {
			var sum int
			for _, item := range items {
				sum += item
			}
			if sum == target {
				res = append(res, items)
			}
			break
		}
	}
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	// 和が target と一致したとき
	// target を nums[i] で引き続けているため、0 になったときが一致となる
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
	}
	for i := index; i < len(nums); i++ {
		// target より大きい場合は除く
		if nums[i] > target {
			break
		}
		c = append(c, nums[i])
		findcombinationSum(nums, target-nums[i], i, c, res)
		// c へ追加した数値を除く
		c = c[:len(c)-1]
	}
}
