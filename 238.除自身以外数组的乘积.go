/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */
package main

// @lc code=start
func productExceptSelf(nums []int) []int {
	l := len(nums)
	var result = make([]int, l)
	t := 1
	for i := 0; i < l; i++ {
		result[i] = t
		t *= nums[i]
	}
	t = 1
	for i := l - 1; i >= 0; i-- {
		result[i] *= t
		t *= nums[i]
	}
	return result
}

// @lc code=end
