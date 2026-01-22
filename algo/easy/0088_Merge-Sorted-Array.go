package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmpNums1 := make([]int, m)
	for i := 0; i < m; i++ {
		tmpNums1[i] = nums1[i]
	}
	for i, j := 0, 0; i+j < m+n; {
		if i == m {
			nums1[i+j] = nums2[j]
			j++
		} else if j == n {
			nums1[i+j] = tmpNums1[i]
			i++
		} else if tmpNums1[i] <= nums2[j] {
			nums1[i+j] = tmpNums1[i]
			i++
		} else {
			nums1[i+j] = nums2[j]
			j++
		}
	}
}
