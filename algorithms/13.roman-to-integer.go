package main

func romanToInt(s string) int {
	romans := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	rs := []rune(s)
	var result int

	f := func(roman string, rs []rune) bool {
		for i, r := range roman {
			if i >= len(rs) {
				return false
			}
			if rs[i] != r {
				return false
			}
		}
		return true
	}

	for {
		if len(rs) == 0 {
			break
		}
		var num int
		var targetRoman string
		for roman, romanNum := range romans {
			if !f(roman, rs) {
				continue
			}
			if romanNum > num {
				num = romanNum
				targetRoman = roman
			}
		}
		result = result + num
		rs = rs[len(targetRoman):]
	}
	return result
}
