package main

func generateParenthesis(n int) []string {
	all := []string{"("}
	for i := 0; i < (n*2 - 1); i++ {
		var tmp []string
		for _, r := range all {
			tmp = append(tmp, r+"(", r+")")
		}
		all = tmp
	}

	var res []string
	for _, a := range all {
		var left []rune
		for _, r := range a {
			if r == '(' {
				left = append(left, r)
			} else {
				if len(left) == 0 {
					left = append(left, '(')
					break
				}
				left = left[:len(left)-1]
			}
		}
		if len(left) > 0 {
			continue
		}
		res = append(res, a)
	}
	return res
}
