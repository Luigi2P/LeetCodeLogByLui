/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		iter1 int = 0
		iter2 int = 0
		r         = make([]int, 0, m+n)
	)
	for {
		if iter1 == m {
			r = append(r, nums2[iter2:]...)
			break
		}
		if iter2 == n {
			r = append(r, nums1[iter1:]...)
			break
		}
		if nums1[iter1] < nums2[iter2] {
			r = append(r, nums1[iter1])
			iter1++
		} else {
			r = append(r, nums2[iter2])
			iter2++
		}
	}
	copy(nums1, r)
	// return nums1
}

// func main() {
// 	var (
// 		nums1 []int = []int{1, 2, 2, 0, 0, 0}
// 		nums2 []int = []int{3, 5, 6}
// 	)
// 	fmt.Println(merge(nums1, 3, nums2, len(nums2)))
// }

// @lc code=end
