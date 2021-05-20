package main

import (
	"math"
)

func divide(dividend int, divisor int) int {
	var intDividend int
	var intDivisor int
	var minus bool
	switch {
	case dividend < 0 && divisor < 0:
		intDividend = 0 - dividend
		intDivisor = 0 - divisor
	case dividend < 0:
		intDividend = 0 - dividend
		intDivisor = divisor
		minus = true
	case divisor < 0:
		intDividend = dividend
		intDivisor = 0 - divisor
		minus = true
	default:
		intDividend = dividend
		intDivisor = divisor
	}
	if intDivisor > intDividend {
		return 0
	}
	count := 1
	tmp := intDivisor
	for {
		if count >= math.MaxInt32 {
			if minus {
				return 0 - math.MaxInt32 - 1
			} else {
				return math.MaxInt32
			}
		}
		if tmp == intDividend {
			if minus {
				return 0 - count
			} else {
				return count
			}
		}
		if tmp > intDividend {
			if minus {
				return 0 - (count - 1)
			} else {
				return count - 1
			}
		}
		tmp = tmp + intDivisor
		count++
	}
}

func divide1(divided int, divisor int) int {
	if divided == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	result := 0
	sign := -1
	if divided > 0 && divisor > 0 || divided < 0 && divisor < 0 {
		sign = 1
	}
	dvd, dvs := abs(divided), abs(divisor)
	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1
			m <<= 1
		}
		dvd -= temp
		result += m
	}
	return sign * result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
