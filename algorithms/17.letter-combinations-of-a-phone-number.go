package main

func letterCombinations(digits string) []string {
	dm := map[string][]string{
		"1": []string{},
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}

	as, bs, cs, ds := []string{}, []string{}, []string{}, []string{}
	if len(digits) == 4 {
		as = dm[string(digits[0])]
		bs = dm[string(digits[1])]
		cs = dm[string(digits[2])]
		ds = dm[string(digits[3])]
	}
	if len(digits) == 3 {
		as = dm[string(digits[0])]
		bs = dm[string(digits[1])]
		cs = dm[string(digits[2])]
	}
	if len(digits) == 2 {
		as = dm[string(digits[0])]
		bs = dm[string(digits[1])]
	}
	if len(digits) == 1 {
		as = dm[string(digits[0])]
	}

	res := []string{}
	for _, a := range as {
		if len(bs) > 0 {
			for _, b := range bs {
				if len(cs) > 0 {
					for _, c := range cs {
						if len(ds) > 0 {
							for _, d := range ds {
								res = append(res, a+b+c+d)
							}
						} else {
							res = append(res, a+b+c)
						}
					}
				} else {
					res = append(res, a+b)
				}
			}
		} else {
			res = append(res, a)
		}
	}
	return res
}
