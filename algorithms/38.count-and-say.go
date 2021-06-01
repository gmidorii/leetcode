package main

import (
	"fmt"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	num := countAndSay(n - 1)

	var tmp []rune
	var res string
	for i, v := range num {
		if i == 0 {
			tmp = []rune{v}
			continue
		}
		if tmp[0] != v {
			res = res + fmt.Sprintf("%v%v", len(tmp), string(tmp[0]))
			tmp = []rune{v}
			continue
		}
		tmp = append(tmp, v)
	}

	res = res + fmt.Sprintf("%v%v", len(tmp), string(tmp[0]))
	return res
}
