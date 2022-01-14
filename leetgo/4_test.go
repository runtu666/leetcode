package main

import (
	"testing"
)

func Test4(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var p1, p2 int
	var l1 = len(nums1)
	var l2 = len(nums2)
	var merge = make([]int, 0, l1+l2)
	if l1 == 0 && l2 == 0 {
		return 0
	}
	for {
		if l1 == 0 {
			merge = nums2
			break
		}
		if l2 == 0 {
			merge = nums1
			break
		}
		if nums2[p2] > nums1[p1] {
			merge = append(merge, nums1[p1])
			if p1 < l1-1 {
				p1++
				continue
			}
			merge = append(merge, nums2[p2:]...)
			break
		} else {
			merge = append(merge, nums2[p2])
			if p2 < l2-1 {
				p2++
				continue
			}
			merge = append(merge, nums1[p1:]...)
			break
		}
	}
	middle := (l1 + l2) / 2
	if (l1+l2)%2 == 0 {
		return (float64(merge[middle]) + float64(merge[middle-1])) / 2
	}
	return float64(merge[middle])
}
