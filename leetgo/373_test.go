package main

import (
	"testing"
)

/**
给定两个以 升序排列 的整数数组 nums1 和 nums2,以及一个整数 k。

定义一对值(u,v)，其中第一个元素来自nums1，第二个元素来自 nums2。

请找到和最小的 k个数对(u1,v1), (u2,v2) ... (uk,vk)。

示例 1:
输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

示例 2:
输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
    [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

示例 3:
输入: nums1 = [1,2], nums2 = [3], k = 3
输出: [1,3],[2,3]
解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test373(t *testing.T) {
	//t.Log(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	//t.Log(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 3))
	t.Log(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	//var p1, p2 int
	var ret = make([][]int, 0, k)
	//var max1 = len(nums1) - 1
	//var max2 = len(nums2) - 1
	//for {
	//	fmt.Println(p1, p2)
	//	ret = append(ret, []int{nums1[p1], nums2[p2]})
	//	if len(ret) == k {
	//		break
	//	}
	//	if p1 == max1 && p2 == max2 {
	//		break
	//	}
	//	if p1 == max1 {
	//		p2++
	//		continue
	//	}
	//	if p2 == max2 {
	//		p1++
	//		continue
	//	}
	//	switch {
	//	case nums1[p1] == nums2[p2]:
	//		if nums2[p2+1] > nums1[p1+1] {
	//			p1++
	//		} else {
	//			p2++
	//		}
	//	case nums1[p1] > nums2[p2]:
	//		p1++
	//	case nums1[p1] < nums2[p2]:
	//		p2++
	//	}
	//}
	return ret
}
