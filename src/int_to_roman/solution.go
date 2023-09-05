package int_to_romain

import "strings"

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
