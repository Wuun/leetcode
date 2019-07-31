package main

func main() {

}

//感觉可以使用传说中的递归或者是函数式编程来做，因为这个解题的过程可以分解为无数个步骤相似的过程。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 {
		if len2%2 == 0 {
			return float64((nums2[len2/2-1] + nums2[len2/2]) / 2)
		}
		return float64(nums2[len2/2])
	} else if len2 == 0 {
		if len1%2 == 0 {
			return float64((nums1[len1/2-1] + nums1[len1/2]))
		}
		return float64(nums1[len1/2])
	}

	if nums1[len1/2] < nums2[len2/2] {
		return findMedianSortedArrays(nums1[len1/2:], nums2[:len2/2])
	} else if nums1[int(len1/2)] > nums2[int(len2/2)] {
		return findMedianSortedArrays(nums1[:len1/2], nums2[len2/2:])
	}
	return float64(nums1[len1/2])
}
