/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */
package main

import (
	"fmt"
)

// @lc code=start

func grayCode(n int) []int {
	var (
		length = 1 << n
		t      = make([]int, 0, length)
	)
	for i := 0; i < length; i++ {
		t = append(t, i^(i/2))
	}
	return t
}

// @lc code=end
func main() {
	fmt.Print(grayCode(3))
}
