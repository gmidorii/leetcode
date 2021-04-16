package main

func maxArea(height []int) int {
	rHeight := make([]int, len(height))
	j := 0
	for i := len(height) - 1; i >= 0; i-- {
		rHeight[j] = height[i]
		j++
	}

	max := 0
	l := len(height)
	for i, v := range height {
		if max > (l-(i+1))*v {
			continue
		}
		for j, rv := range rHeight {
			if i+j >= l {
				break
			}
			w := (l - j) - (i + 1)
			h := 0
			if rv > v {
				h = v
			} else {
				h = rv
			}
			area := w * h
			if area > max {
				max = area
			}
		}
	}

	return max
}
