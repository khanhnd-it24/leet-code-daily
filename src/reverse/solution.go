package reverse

import (
	"math"
	"strconv"
)

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the
signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21

Constraints:
	- -2^31 <= x <= 2^31 - 1
*/

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}
	str := strconv.Itoa(x)
	strReverse := reverseStr(str)
	x, _ = strconv.Atoi(strReverse)
	if isNegative {
		x = -x
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	return x
}

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse2(x int) int {
	r := 0
	for x != 0 {
		r = r*10 + x%10
		x /= 10
	}
	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return r
}
