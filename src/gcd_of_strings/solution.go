package gcd_of_strings

import "strings"

/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t
(i.e., t is concatenated with itself one or more times).
Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Example 2:
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"

Example 3:
Input: str1 = "LEET", str2 = "CODE"
Output: ""

Constraints:

 - 1 <= str1.length, str2.length <= 1000
 - str1 and str2 consist of English uppercase letters.
*/

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)
	gcdLen := gcd(len1, len2)
	n1, n2 := len1/gcdLen, len2/gcdLen
	if strings.Repeat(str1[:gcdLen], n1) != str1 || strings.Repeat(str2[:gcdLen], n2) != str2 {
		return ""
	}
	if str1[:gcdLen] == str2[:gcdLen] {
		return str1[:gcdLen]
	}
	return ""
}

func gcdOfStrings2(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
