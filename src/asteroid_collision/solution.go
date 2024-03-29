package asteroid_collision

/*
We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction
(positive meaning right, negative meaning left). Each asteroid moves at the same speed.
Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode.
If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.



Example 1:
Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

Example 2:
Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.

Example 3:
Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

Constraints:

 - 2 <= asteroids.length <= 104
 - -1000 <= asteroids[i] <= 1000
 - asteroids[i] != 0
*/

func asteroidCollision(asteroids []int) []int {
	res := make([]int, 0)
	i := 0
	for i < len(asteroids) {
		if !isNegative(asteroids[i]) {
			res = append(res, asteroids[i])
			i++
			continue
		}
		if len(res) == 0 || isNegative(res[len(res)-1]) {
			res = append(res, asteroids[i])
			i++
			continue
		}
		switch compareAbs(res[len(res)-1], asteroids[i]) {
		case 1:
			i++
			break
		case -1:
			res = res[:len(res)-1]
			break
		default:
			res = res[:len(res)-1]
			i++
		}
	}
	return res
}

func isNegative(a int) bool {
	return a < 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func compareAbs(a, b int) int {
	if abs(a) > abs(b) {
		return 1
	}
	if abs(a) < abs(b) {
		return -1
	}
	return 0
}
