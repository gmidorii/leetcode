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

func maxArea_2(height []int) int {
	leftIdx := 0
	rightIdx := len(height) - 1
	max := 0
	for {
		if leftIdx == rightIdx {
			break
		}
		lh := height[leftIdx]
		rh := height[rightIdx]
		w := rightIdx - leftIdx
		if lh > rh {
			area := rh * w
			if area > max {
				max = area
			}
			rightIdx = rightIdx - 1
			continue
		} else {
			area := lh * w
			if area > max {
				max = area
			}
			leftIdx = leftIdx + 1
			continue
		}
	}
	return max
}