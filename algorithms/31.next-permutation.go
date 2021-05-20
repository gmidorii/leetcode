package main

import "fmt"

var l map[int]string = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

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
