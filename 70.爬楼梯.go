/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// package main

func climbStairs(n int) int {
	var (
		t  = 0
		p  = 1
		pp = 1
	)
	if n < 2 {
		return 1
	}
	for i := 2; i <= n; i++ {
		t = p + pp
		pp = p
		p = t
	}
	return t
}

// @lc code=end
