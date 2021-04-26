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

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)

func letterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
}
