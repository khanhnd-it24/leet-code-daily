package long_common_prefix

/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
	Input: strs = ["flower","flow","flight"]
	Output: "fl"

Example 2:
	Input: strs = ["dog","racecar","car"]
	Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:

 - 1 <= strs.length <= 200
 - 0 <= strs[i].length <= 200
 - strs[i] consists of only lowercase English letters.
*/

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
