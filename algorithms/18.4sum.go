package main

import "sort"

func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4)
}

func kSum(nums []int, target, k int) [][]int {
	res := [][]int{}
	if len(nums) == 0 || nums[0]*k > target || target > nums[len(nums)-1]*k {
		return res
	}
	if k == 2 {
		return twoSum(nums, target)
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			for _, set := range kSum(nums[i+1:], target-nums[i], k-1) {
				res = append(res, append(set, nums[i]))
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	res := [][]int{}
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		switch {
		case sum < target || (lo > 0 && nums[lo] == nums[lo-1]):
			lo = lo + 1
		case sum > target || (hi < len(nums)-1 && nums[hi] == nums[hi+1]):
			hi = hi - 1
		default:
			res = append(res, []int{nums[lo], nums[hi]})
			lo = lo + 1
			hi = hi - 1
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for m := k + 1; m < len(nums); m++ {
					if (nums[i] + nums[j] + nums[k] + nums[m]) == target {
						res = append(res, []int{nums[i], nums[j], nums[k], nums[m]})
					}
				}
			}
		}
	}

	return distinct(res)
}

func distinct(src [][]int) [][]int {
	ts := make([][]int, len(src))
	copy(ts, src)
	ngList := make([]bool, len(src))
	for i, s := range src {
		if ngList[i] {
			continue
		}
		for j, t := range ts {
			if i == j {
				continue
			}
			if matchNoOrder(s, t) {
				ngList[j] = true
			}
		}
	}
	var res [][]int
	for i, ng := range ngList {
		if !ng {
			res = append(res, src[i])
		}
	}
	return res
}

func matchNoOrder(ss []int, tt []int) bool {
	if len(ss) != len(tt) {
		return false
	}
	var tmp []int
	for _, s := range ss {
		m := false
		for _, t := range tt {
			if !m && s == t {
				m = true
				continue
			}
			if m || s != t {
				tmp = append(tmp, t)
			}
		}
		tt = tmp
		tmp = []int{}
	}
	return len(tt) == 0
}
