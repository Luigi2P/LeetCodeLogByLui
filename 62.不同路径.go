/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
package main

// import "fmt"

func uniquePaths(m int, n int) int {
	var (
		a = m + n - 2
		b int
		r = 1
		i int
	)
	if m > n {
		b = n - 1
	} else {
		b = m - 1
	}
	for i = 0; i < b; i++ {
		r *= a - i
	}
	for i = b; i >= 1; i-- {
		r /= i
	}
	return r
}

// @lc code=end
// func main() {
// 	fmt.Println(uniquePaths(23, 12))
// }
