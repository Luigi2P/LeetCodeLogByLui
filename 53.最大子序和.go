/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
package main

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	max := nums[0]
	for i := 1; i < l; i++ {
		if nums[i] < nums[i-1]+nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// func main() {
// 	t := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// 	println(maxSubArray(t))
// }

// @lc code=end
