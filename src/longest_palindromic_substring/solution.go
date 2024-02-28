package longest_palindromic_substring

import "strings"

/*
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Constraints:
 - 1 <= s.length <= 1000
 - s consist of only digits and English letters.
*/

type palindrome struct {
	value string
	start int
	end   int
}

func (p palindrome) len() int {
	return p.end - p.start + 1
}

func longestPalindrome(s string) string {
	if isPalindrome(s) {
		return s
	}

	chars := strings.Split(s, "")
	lenDP := len(chars) / 2
	lenChar := len(chars)
	dp := make([][]palindrome, lenDP)
	palindromes := make([]palindrome, 0)
	for i, char := range chars {
		palindromes = append(palindromes, palindrome{value: char, start: i, end: i})
		if i != lenChar-1 && chars[i] == chars[i+1] {
			palindromes = append(palindromes, palindrome{value: char + chars[i+1], start: i, end: i + 1})
		}
	}
	dp[0] = palindromes

	for i := 1; i < lenDP; i++ {
		oldPalindromes := dp[i-1]
		newPalindromes := make([]palindrome, len(oldPalindromes))
		for j := 0; j < len(oldPalindromes); j++ {
			palindromeJ := oldPalindromes[j]
			if palindromeJ.start == 0 || palindromeJ.end == lenChar-1 || chars[palindromeJ.start-1] != chars[palindromeJ.end+1] {
				newPalindromes[j] = palindromeJ
			} else {
				newPalindromes[j] = palindrome{
					value: chars[palindromeJ.start-1] + palindromeJ.value + chars[palindromeJ.end+1],
					start: palindromeJ.start - 1,
					end:   palindromeJ.end + 1,
				}
			}
		}
		dp[i] = newPalindromes
	}
	longest := dp[lenDP-1][0]
	for i := 1; i < len(dp[lenDP-1]); i++ {
		if longest.len() < dp[lenDP-1][i].len() {
			longest = dp[lenDP-1][i]
		}
	}
	return longest.value
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome2(s string) string {
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)
	P := make([]int, n)
	C, R := 0, 0

	for i := 1; i < n-1; i++ {
		if R > i {
			P[i] = min(R-i, P[2*C-i])
		}
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}
		if i+P[i] > R {
			C, R = i, i+P[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}
	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
