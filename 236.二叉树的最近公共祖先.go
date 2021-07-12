/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) int
	var result *TreeNode = nil
	dfs = func(node *TreeNode) int {
		r := 0
		if node == nil {
			return 0
		}
		if r == -1 {
			return -1
		}
		if node == p || node == q {
			r++
		}
		r += dfs(node.Left)
		r += dfs(node.Right)
		if r == 2 {
			result = node
			r = -1
		}
		return r
	}
	dfs(root)
	return result
}

// @lc code=end
