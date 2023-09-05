package long_common_prefix

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func commonPrefix(left, right string) string {
	_min := min(len(left), len(right))
	for i := 0; i < _min; i++ {
		if left[i] != right[i] {
			return left[0:i]
		}
	}
	return left[0:_min]
}

func longestCommonPrefixBinary(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}

	mid := (l + r) / 2
	lcpLeft := longestCommonPrefixBinary(strs, l, mid)
	lcpRight := longestCommonPrefixBinary(strs, mid+1, r)
	return commonPrefix(lcpLeft, lcpRight)
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	return longestCommonPrefixBinary(strs, 0, len(strs)-1)
}
