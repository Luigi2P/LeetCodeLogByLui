/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	max := 0
	t := 0
	for _, i := range prices {
		t = i - min
		if t > max {
			max = t
		}
		if i < min {
			min = i
		}
	}
	return max
}

// @lc code=end
