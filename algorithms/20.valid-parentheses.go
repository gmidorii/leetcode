package main

func isValid(s string) bool {
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []rune
	for _, r := range s {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, r)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		got := stack[len(stack)-1]
		want := m[r]
		if got != want {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
