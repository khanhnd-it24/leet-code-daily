package buy_and_sell_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.

Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.

Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
*/

func maxProfit(prices []int) int {
	profit, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			profit += prices[i-1] - min
			min = prices[i]
		}
		if min > prices[i] {
			min = prices[i]
		}
		if i == len(prices)-1 && prices[i] > min {
			profit += prices[i] - min
		}
	}
	return profit
}

var testcase = []domains.Testcase{
	{
		In:  []int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9},
		Out: 25,
	},
	{
		In:  []int{1, 2, 3, 4, 5},
		Out: 4,
	},
	{
		In:  []int{7, 6, 4, 3, 1},
		Out: 0,
	},
}

func TestBuyAndSellII(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, maxProfit(input))
	}
}
