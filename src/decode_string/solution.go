package decode_string

import "strconv"

/*
Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly
k times. Note that k is guaranteed to be a positive integer.
You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed,
etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those
repeat numbers, k. For example, there will not be input like 3a or 2[4].
The test cases are generated so that the length of the output will never exceed 105.

Example 1:
Input: s = "3[a]2[bc]"
Output: "aaabcbc"

Example 2:
Input: s = "3[a2[c]]"
Output: "accaccacc"

Example 3:
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"

Constraints:
 - 1 <= s.length <= 30
 - s consists of lowercase English letters, digits, and square brackets '[]'.
 - s is guaranteed to be a valid input.
 - All the integers in s are in the range [1, 300].
*/

func decodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] == '[' {
					a := j - 1
					for x := j - 1; x >= 0; x-- {
						if stack[x] < '0' || stack[x] > '9' {
							a = x + 1
							break
						}
						if x == 0 {
							a = x
						}
					}
					n, _ := strconv.Atoi(string(stack[a:j]))
					currentLen := len(stack)
					for k := 0; k < n-1; k++ {
						stack = append(stack, stack[j+1:currentLen]...)
					}
					stack = append(stack[:a], stack[j+1:]...)
					break
				}
			}
		}
	}
	return string(stack)
}
