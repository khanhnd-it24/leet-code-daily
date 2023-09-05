package check_if_exist

/*
Given an array arr of integers, check if there exist two indices i and j such that :

 - i != j
 - 0 <= i, j < arr.length
 - arr[i] == 2 * arr[j]

Example 1:
	Input: arr = [10,2,5,3]
	Output: true
	Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

Example 2:
	Input: arr = [3,1,7,11]
	Output: false
	Explanation: There is no i and j that satisfy the conditions.

Constraints:
 - 2 <= arr.length <= 500
 - -103 <= arr[i] <= 103
*/

func checkIfExist(arr []int) bool {
	hashTable := map[int]int{}
	for i, v := range arr {
		hashTable[v*2] = i
	}

	for i, v := range arr {
		j, ok := hashTable[v]
		if ok && i != j {
			return true
		}
	}

	return false
}
