package reverse_vowels

/*
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:
Input: s = "hello"
Output: "holle"

Example 2:
Input: s = "leetcode"
Output: "leotcede"

Constraints:
 - 1 <= s.length <= 3 * 105
 - s consist of printable ASCII characters.
*/

func reverseVowels(s string) string {
	var vowels []int
	if len(s) == 1 {
		return s
	}

	for i, ss := range s {
		if ss == 'a' || ss == 'e' || ss == 'i' || ss == 'o' || ss == 'u' ||
			ss == 'A' || ss == 'E' || ss == 'I' || ss == 'O' || ss == 'U' {
			vowels = append(vowels, i)
		}
	}

	runes := []rune(s)
	for i := 0; i < len(vowels)/2; i++ {
		runes[vowels[i]], runes[vowels[len(vowels)-i-1]] = runes[vowels[len(vowels)-i-1]], runes[vowels[i]]
	}

	return string(runes)
}
