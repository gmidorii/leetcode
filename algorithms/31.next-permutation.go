package main

import "fmt"

/*
1. 末尾から順番に比較して若いほうが小さい値の場合は入れ替えて返却
2. 先頭との比較で先頭のほうが小さい場合は、先頭の値より最も近く大きい値と入れ替えて、小さい順にソート
3. 先頭との比較で先頭のほうが大きい場合は、逆順にして返却
*/

func nextPermutation(nums []int) {
	lastIdx := len(nums) - 1
	for {
		if lastIdx <= 0 {
			break
		}
		for i := 0; i < lastIdx; i++ {
			if lexiCompare(nums[i], nums[i+1]) == -1 {
				nums[i], nums[i+1] = swap(nums[i], nums[i+1])
				fmt.Println(swap(nums[i], nums[i+1]))
			}
		}
		lastIdx--
	}
}

func lexiCompare(a, b int) int {
	// 1, 0, -1
	sa := l[a]
	sb := l[b]
	switch {
	case sa > sb:
		return -1
	case sa == sb:
		return 0
	default:
		return 1
	}
}

func swap(a, b int) (int, int) {
	return b, a
}
