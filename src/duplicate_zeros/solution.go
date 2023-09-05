package duplicate_zeros

// 1, 0, 2, 3, 0, 4, 5, 0
// 1, 0, 0, 2, 3, 0, 0, 4
func duplicateZeros(arr []int) {
	length := len(arr)

	result := make([]int, length)
	for i, j := 0, 0; j < length; i, j = i+1, j+1 {
		result[j] = arr[i]
		if result[j] == 0 {
			j++
		}
	}
	for i := 0; i < length; i++ {
		arr[i] = result[i]
	}
}
