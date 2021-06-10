package main

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	findCombinations2(candidates, 0, target, c, &res)
	return res
}

func findCombinations2(candidates []int, index, target int, c []int, res *[][]int) {
	if target == 0 {
		if duplicate(c, *res) {
			return
		}
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		c = append(c, candidates[i])
		findCombinations2(candidates, i+1, target-candidates[i], c, res)
		c = c[:len(c)-1]
	}
}

func duplicate(targets []int, res [][]int) bool {
	for _, rs := range res {
		if matchSlice2(rs, targets) {
			return true
		}
	}
	return false
}

func matchSlice2(ss, targets []int) bool {
	if len(ss) != len(targets) {
		return false
	}
	matchCount := 0
	tmp := make([]int, len(targets))
	copy(tmp, targets)
	for _, s := range ss {
		for i, t := range tmp {
			if s == t {
				matchCount++
				tmp[i] = -1
				break
			}
		}
	}
	if matchCount == len(targets) {
		return true
	}
	return false
}
