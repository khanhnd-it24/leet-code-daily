package check_if_exist

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
