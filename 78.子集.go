/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
package main

// import "fmt"

func subsets(nums []int) [][]int {
	var (
		r      = make([][]int, 0)
		t      = make([]int, 0)
		length = len(nums)
		dfs    func(int)
	)
	dfs = func(cur int) {
		if cur == length {
			r = append(r, append([]int(nil), t...))
			return
		}
		t = append(t, nums[cur])
		dfs(cur + 1)
		t = t[:len(t)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return r
}

// func main() {
// 	t := []int{1, 2, 3}
// 	fmt.Println(subsets(t))
// }

// @lc code=end
