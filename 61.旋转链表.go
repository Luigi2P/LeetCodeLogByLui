/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func rotateRight(head *ListNode, k int) *ListNode {
	iter := head
	length := 1
	if head == nil || k == 0 {
		return head
	}
	for iter.Next != nil {
		iter = iter.Next
		length++
	}
	iter.Next = head
	prev := iter
	m := length - k%length
	for i := 0; i <= m; i++ {
		prev = iter
		iter = iter.Next
	}
	prev.Next = nil
	return iter
}

// @lc code=end
