package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return 0
	}
	var idx int
	var haystackFirstIdx int
	var needleIdx int
	for {
		if needleIdx >= len(needle) {
			return haystackFirstIdx
		}
		if haystack[idx] == needle[needleIdx] {
			idx++
			needleIdx++
		} else {
			haystackFirstIdx++
			idx = haystackFirstIdx
			needleIdx = 0
		}
	}
	return 0
}
