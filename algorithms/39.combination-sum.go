package main

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
