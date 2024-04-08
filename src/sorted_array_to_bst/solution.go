package sorted_array_to_bst

/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced binary search tree.

Example 1:

Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.


Constraints:
 - 1 <= nums.length <= 104
 - -10^4 <= nums[i] <= 10^4
 - nums is sorted in a strictly increasing order.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return toBTS(nums, 0, len(nums)-1)
}

func toBTS(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  toBTS(nums, l, mid-1),
		Right: toBTS(nums, mid+1, r),
	}

	return root
}
