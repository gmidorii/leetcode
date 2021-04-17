package main

func intToRoman(num int) string {
	v := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	romanInts := []int{1000, 100, 10, 1}
	var result string
	for _, ri := range romanInts {
		r := num / ri
		roman := v[ri]
		middle, _ := v[ri*5]
		pRoman, _ := v[ri*10]
		var x string
		for i := 1; i <= r; i++ {
			if i == 4 {
				x = roman + middle
				continue
			}
			if i == 5 {
				x = middle
				continue
			}
			if i == 9 {
				x = roman + pRoman
				continue
			}
			x = x + roman
		}

		result = result + x
		num = num % ri
	}

	return result
}
