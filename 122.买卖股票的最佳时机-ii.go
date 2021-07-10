/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	total := 0
	prev := prices[0]
	var t int
	for _, i := range prices {
		t = i - prev
		if t > 0 {
			total += t
		}
		prev = i
	}
	return total
}

// @lc code=end
