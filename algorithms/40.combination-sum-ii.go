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
