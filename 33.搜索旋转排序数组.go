/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
// package leetcode

func search(nums []int, target int) int {
	return findR(nums, target, 9, len(nums)-1)
}

func findR(nums []int, target int, left int, right int) int {
	mid := (left + right) / 2
	if left == right {
		if nums[mid] == target {
			return mid
		} else {
			return -1
		}
	}
	if nums[left] <= nums[mid] && target >= nums[left] {
		return findR(nums, target, left, mid)
	} else {
		return findR(nums, target, mid+1, right)
	}
}

// @lc code=end
