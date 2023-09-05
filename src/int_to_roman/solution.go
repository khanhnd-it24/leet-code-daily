package int_to_romain

import "strings"

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply
X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

 - I can be placed before V (5) and X (10) to make 4 and 9.
 - X can be placed before L (50) and C (100) to make 40 and 90.
 - C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.

Example 1:
	Input: num = 3
	Output: "III"
	Explanation: 3 is represented as 3 ones.

Example 2:
	Input: num = 58
	Output: "LVIII"
	Explanation: L = 50, V = 5, III = 3.

Example 3:
	Input: num = 1994
	Output: "MCMXCIV"
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Constraints:
 - 1 <= num <= 3999
*/

var orders = [4]map[int]byte{
	{
		1: 'I',
		5: 'V',
	},
	{
		1: 'X',
		5: 'L',
	},
	{
		1: 'C',
		5: 'D',
	},
	{
		1: 'M',
	},
}

func intToRoman(num int) string {
	builder := &strings.Builder{}
	var res []string
	order := 0

	for num > 0 {
		digit := num % 10
		num = num / 10
		digitToRoman(digit, order, builder)
		order++

		res = append(res, builder.String())
		builder.Reset()
	}

	for i := len(res) - 1; i >= 0; i-- {
		builder.WriteString(res[i])
	}

	return builder.String()
}

func digitToRoman(digit int, order int, builder *strings.Builder) {
	if digit == 4 {
		builder.WriteByte(orders[order][1])
		digit = 5
	}

	if digit == 9 {
		builder.WriteByte(orders[order][1])
		digit = 1
		order = order + 1
	}

	if digit >= 5 {
		builder.WriteByte(orders[order][5])
		digit -= 5
	}

	for digit > 0 {
		digit--
		builder.WriteByte(orders[order][1])
	}
}
