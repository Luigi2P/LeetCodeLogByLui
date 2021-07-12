/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

package main

// @lc code=start
func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	k = k - 1
	var i int
	for {
		i = partition(nums, left, right)
		if i > k {
			right = i - 1
		} else if i < k {
			left = i + 1
		} else if i == k {
			return nums[i]
		}
	}
}
func partition(nums []int, left int, right int) int {
	if left > right {
		return -1
	}
	i := left
	j := right
	t := nums[left]
	for i < j {
		for i < j && nums[j] <= t {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] >= t {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = t
	return i
}

// @lc code=end
