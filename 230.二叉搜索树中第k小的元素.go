/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(node *TreeNode, k int) int {
	var (
		dfs  func(node *TreeNode)
		nums = make([]int, 0, k)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if len(nums) == k {
			return
		}
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
		return
	}
	dfs(node)
	return nums[k-1]
}

// @lc code=end
