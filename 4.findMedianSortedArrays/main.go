package main

func main() {

}

//we may solve this problem in the way of recurse ot FP way ,for break this problem to lots of pieces.
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
			return float64((nums1[len1/2-1] + nums1[len1/2]) / 2)
		}
		return float64(nums1[len1/2])
	}

	if len1%2 == 0 && len2%2 == 0 {
		if nums1[len1/2-1] <= nums2[len2/2-1] {
			return findMedianSortedArrays(nums1[len1/2:], nums2[:len2/2])
		} else if nums1[len1/2-1] > nums2[len2/2-1] {
			return findMedianSortedArrays(nums1[:len1/2+1], nums2[len2/2:])
		}
	} else if len1%2 == 0 && len2%2 == 1 {
		if nums1[len1/2-1] <= nums2[len2/2] {
			return findMedianSortedArrays(nums1[len1/2:], nums2[:len2/2+1])
		} else if nums1[len1/2-1] > nums2[len2/2] {
			return findMedianSortedArrays(nums1[:len1/2], nums2[len2/2:])
		}
	} else if len1%2 == 1 && len2%2 == 0 {
		if nums1[len1/2] <= nums2[len2/2-1] {
			return findMedianSortedArrays(nums1[len1/2:], nums2[:len2/2])
		} else if nums1[len1/2] > nums2[len2/2-1] {
			return findMedianSortedArrays(nums1[:len1/2+1], nums2[len2/2:])
		}
	}

	if nums1[len1/2] <= nums2[len2/2] {
		return findMedianSortedArrays(nums1[len1/2:], nums2[:len2/2+1])
	}

	return findMedianSortedArrays(nums1[:len1/2+1], nums2[len2/2:])
}

func secondChallenge(nums1 []int, nums2 []int) float64 {
	return 0
}
