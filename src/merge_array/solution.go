package merge_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p3 := m-1, n-1, m+n-1
	for ; p1 >= 0 && p2 >= 0; p3-- {
		if nums1[p1] <= nums2[p2] {
			nums1[p3] = nums2[p2]
			p2--
		} else {
			nums1[p3] = nums1[p1]
			p1--
		}
	}

	if p2 >= 0 {
		copy(nums1[:p3+1], nums2[:p2+1])
	}
}
