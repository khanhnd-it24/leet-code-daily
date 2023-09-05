package word_pattern

import "strings"

/*
Given a pattern and a string s, find if s follows the same pattern.
Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

Example 1:
	Input: pattern = "abba", s = "dog cat cat dog"
	Output: true

Example 2:
	Input: pattern = "abba", s = "dog cat cat fish"
	Output: false

Example 3:
	Input: pattern = "aaaa", s = "dog cat cat dog"
	Output: false

Constraints:
 - 1 <= pattern.length <= 300
 - pattern contains only lower-case English letters.
 - 1 <= s.length <= 3000
 - s contains only lowercase English letters and spaces ' '.
 - s does not contain any leading or trailing spaces.
 - All the words in s are separated by a single space.
*/

func wordPattern(pattern string, s string) bool {
	sWords := strings.Split(s, " ")
	sLen := len(sWords)
	pLen := len(pattern)
	if sLen != pLen {
		return false
	}

	sMap := make(map[byte]string)
	pMap := make(map[string]byte)

	for i := 0; i < sLen; i++ {
		sWord := sWords[i]
		pChar := pattern[i]
		if val, ok := sMap[pChar]; ok {
			if val != sWord {
				return false
			}
		} else {
			if _, ok := pMap[sWord]; ok {
				return false
			}
			sMap[pChar] = sWord
			pMap[sWord] = pChar
		}
	}
	return true
}
