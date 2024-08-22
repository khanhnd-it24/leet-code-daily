package reconstruct_matrix

// Description: https://leetcode.com/problems/reconstruct-a-2-row-binary-matrix/description

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var emptyRes [][]int

	upperRows := make([]int, len(colsum))
	lowerRows := make([]int, len(colsum))

	countUpper, countLower := 0, 0
	for i, v := range colsum {
		if countLower > lower || countUpper > upper {
			return emptyRes
		}
		if v == 2 {
			upperRows[i] = 1
			lowerRows[i] = 1
			countUpper++
			countLower++
		} else if v == 0 {
			upperRows[i] = 0
			lowerRows[i] = 0
		}
	}
	for i, v := range colsum {
		if countLower > lower || countUpper > upper {
			return emptyRes
		}
		if v == 1 {
			if countUpper < upper {
				upperRows[i] = 1
				lowerRows[i] = 0
				countUpper++
			} else {
				upperRows[i] = 0
				lowerRows[i] = 1
				countLower++
			}
		}
	}
	if countLower != lower || countUpper != upper {
		return emptyRes
	}

	return [][]int{upperRows, lowerRows}
}
