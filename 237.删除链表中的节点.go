/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	var prev *ListNode
	for node.Next != nil {
		node.Val = node.Next.Val
		prev = node
		node = node.Next
	}
	prev.Next = nil
}

// @lc code=end
