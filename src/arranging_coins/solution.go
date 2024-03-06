package arranging_coins

import "math"

/*
You have n coins and you want to build a staircase with these coins.
The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.
Given the integer n, return the number of complete rows of the staircase you will build.

Example 1:
Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.

Example 2:
Input: n = 8
Output: 3
Explanation: Because the 4th row is incomplete, we return 3.


Constraints:
 - 1 <= n <= 231 - 1
*/
// (k*(k+1))/2 <= n => k^2 < 2n => k < sqrt(2n)
func arrangeCoins(n int) int {
	a := int(math.Sqrt(float64(2 * n)))
	for i := a; i >= 1; i-- {
		if (i*(i+1))/2 <= n {
			return i
		}
	}
	return 1
}
