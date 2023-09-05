package valid_mountain_array

func validMountainArray(arr []int) bool {
	len := len(arr)
	if len <= 1 || arr[0] >= arr[1] {
		return false
	}

	isMountainArray := false
	for i := 2; i < len; i++ {
		if arr[i] < arr[i-1] {
			isMountainArray = true
		} else if arr[i] == arr[i-1] || (isMountainArray && arr[i] > arr[i-1]) {
			return false
		} else {
			isMountainArray = false
		}
	}

	return isMountainArray
}
