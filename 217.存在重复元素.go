/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */
package main

// @lc code=start
func containsDuplicate(nums []int) bool {
	var set = make(map[int]struct{})
	for _, v := range nums {
		_, ok := set[v]
		if ok == true {
			return true
		} else {
			set[v] = struct{}{}
		}
	}
	return false
}

// @lc code=end
