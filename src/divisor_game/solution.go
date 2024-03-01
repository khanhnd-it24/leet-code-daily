package divisor_game

/*
Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number n on the chalkboard. On each player's turn, that player makes a move consisting of:

Choosing any x with 0 < x < n and n % x == 0.
Replacing the number n on the chalkboard with n - x.
Also, if a player cannot make a move, they lose the game.

Return true if and only if Alice wins the game, assuming both players play optimally.



Example 1:

Input: n = 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.
Example 2:

Input: n = 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.


Constraints:

1 <= n <= 1000
*/

func divisorGame(n int) bool {
	if n <= 1 {
		return false
	}
	dp := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = false
	}
	dp[2] = true
	for i := 3; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
			}
		}
	}
	return dp[n]
}
