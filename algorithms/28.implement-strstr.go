package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	var idx int
	var haystackFirstIdx int
	var needleIdx int
	for {
		if needleIdx >= len(needle) {
			return haystackFirstIdx
		}
		if haystackFirstIdx >= len(haystack) {
			return -1
		}
		if idx >= len(haystack) {
			haystackFirstIdx++
			idx = haystackFirstIdx
			needleIdx = 0
			continue
		}
		if haystack[idx] == needle[needleIdx] {
			idx++
			needleIdx++
		} else {
			haystackFirstIdx++
			idx = haystackFirstIdx
			needleIdx = 0
			continue
		}
	}
}
