package permutation_sequence

import "strconv"

/*
The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Example 1:
Input: n = 3, k = 3
Output: "213"

Example 2:
Input: n = 4, k = 9
Output: "2314"

Example 3:
Input: n = 3, k = 1
Output: "123"
*/

func getPermutation(n int, k int) string {
	numbers := make([]string, n)
	for i := 0; i < n; i++ {
		numbers[i] = strconv.Itoa(i + 1)
	}

	factorials := make([]int, n+1)
	factorials[0] = 1
	for i := 1; i <= n; i++ {
		factorials[i] = factorials[i-1] * i
	}
	k--
	result := ""
	for n > 0 {
		idx := k / factorials[n-1]
		result += numbers[idx]
		numbers = remove(numbers, idx)
		k %= factorials[n-1]
		n--
	}
	return result
}

func remove(slice []string, s int) []string {
	if s == len(slice)-1 {
		return slice[:s]
	}
	return append(slice[:s], slice[s+1:]...)
}
