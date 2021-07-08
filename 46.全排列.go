/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
package main

var r [][]int
var flags [6]bool

func permute(nums []int) [][]int {
	l := len(nums)
	r = make([][]int, 0)
	t := make([]int, 0, 6)
	dfs(nums, l, t)
	return r
}

func dfs(nums []int, left int, t []int) { //返回当前状态下所有可能的排列
	l := len(nums)
	if left == 0 {
		T := make([]int, l, 6)
		copy(T, t)
		r = append(r, T)
		return
	}
	for i := 0; i < l; i++ {
		if flags[i] == false {
			flags[i] = true
			dfs(nums, left-1, append(t, nums[i]))
			flags[i] = false
		}
	}
	return
}

// func main() {
// 	var t = []int{1, 0}
// 	fmt.Println(permute(t))
// }

// @lc code=end
