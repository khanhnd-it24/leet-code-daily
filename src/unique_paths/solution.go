package unique_paths

/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or
right at any point in time.
Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
The test cases are generated so that the answer will be less than or equal to 2 * 109.

Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down


Constraints:
 - 1 <= m, n <= 100
*/

func uniquePaths(m int, n int) int {
	k := m - 1
	x := m - 1 + n - 1
	return c2(x, k)
}

func c(n, k int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		aa := make([]int, n+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				aa[j] = 1
			} else {
				aa[j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
		dp[i] = aa
	}

	return dp[n][k]
}

func c2(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	if k == 1 {
		return n
	}
	return c2(n-1, k) + c2(n-1, k-1)
}

func uniquePathsWithMath(m int, n int) int {
	ans := 1
	for i := 1; i <= m-1; i++ {
		ans = ans * (n - 1 + i) / i
	}
	return ans
}
