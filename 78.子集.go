/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
package main

// import "fmt"

var r = make([][]int, 0)

func subsets(nums []int) [][]int {
	l := len(nums)
	r = make([][]int, 0)
	for i := 0; i <= l; i++ {
		t := make([]int, 0, 10)
		dfs(nums, i, t)
	}
	return r
}
func dfs(nums []int, left int, t []int) {
	length := len(nums)
	if left == 0 {
		r = append(r, append([]int(nil), t...))
		return
	}
	if length < left {
		return
	}
	for i := 0; i < length; i++ {
		dfs(nums[i+1:], left-1, append(t, nums[i]))
	}
	return
}

// func main() {
// 	t := []int{1, 2, 3}
// 	fmt.Println(subsets(t))
// }

// @lc code=end
