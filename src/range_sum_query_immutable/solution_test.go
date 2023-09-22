package range_sum_query_immutable

/*
Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive
(i.e. nums[left] + nums[left + 1] + ... + nums[right]).

Example 1:

Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3


Constraints:

1 <= nums.length <= 10^4
-10^5 <= nums[i] <= 10^5
0 <= left <= right < nums.length
At most 10^4 calls will be made to sumRange.
*/

type NumArray struct {
	size  int
	table []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{}
	numArray.size = len(nums)
	numArray.table = make([]int, numArray.size)
	numArray.table[0] = nums[0]
	for i := 1; i < numArray.size; i++ {
		numArray.table[i] = numArray.table[i-1] + nums[i]
	}
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.table[right]
	}
	return this.table[right] - this.table[left-1]
}
