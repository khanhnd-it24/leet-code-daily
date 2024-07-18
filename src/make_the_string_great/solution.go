package make_the_string_great

// Description: https://leetcode.com/problems/make-the-string-great
func makeGood(s string) string {
	var i int
	for i < len(s)-1 {
		if s[i] == s[i+1]-('a'-'A') || s[i] == s[i+1]+('a'-'A') {
			s = s[:i] + s[i+2:]
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return s
}

func makeGood2(s string) string {
	stack := make([]byte, 0, len(s))
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && diff(stack[len(stack)-1], s[i]) == 32 {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}

func diff(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}
