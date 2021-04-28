package main

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
