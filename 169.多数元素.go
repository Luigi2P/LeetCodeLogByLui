/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package main

// @lc code=start

func majorityElement(nums []int) int {
	var candidate int
	count := 0
	for _, val := range nums {
		if count == 0 {
			candidate = val
		}
		if val == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// @lc code=end
