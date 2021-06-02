package main

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(candidates); i++ {
		c := candidates[i]
		if c > target {
			continue
		}
		if c == target {
			res = append(res, []int{c})
			continue
		}

		// 足して割れるかを順に計算する
		// a+b, a*2+b, a+b*2 で割って 0 になれば良い
		// a+b+c, a*2+b+c, a+b*2+c, a+b+c*2 ...
		// 組み合わせをどうやって絞るか検討
	}
	return res
}
