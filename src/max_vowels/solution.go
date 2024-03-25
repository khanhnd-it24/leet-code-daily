package max_vowels

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

Example 1:
Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.

Example 2:
Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.

Example 3:
Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

Constraints:
 - 1 <= s.length <= 105
 - s consists of lowercase English letters.
 - 1 <= k <= s.length
*/

func maxVowels(s string, k int) int {
	l, sum, maxNow := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if isVowels(s[i]) {
			if i-l+1 < k {
				sum++
				continue
			}
			sum++
		}
		if i-l+1 > k {
			if isVowels(s[l]) {
				sum--
			}
			l++
		}
		if sum > maxNow {
			maxNow = sum
		}
	}
	return maxNow
}

func isVowels(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
