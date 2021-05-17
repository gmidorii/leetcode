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
