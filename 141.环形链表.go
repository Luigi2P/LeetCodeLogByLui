/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
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
func hasCycle(head *ListNode) bool {
	var (
		fast *ListNode = head
		slow *ListNode = head
	)
	if head == nil {
		return false
	}
	for {
		if fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return false
		}
		if fast == slow {
			return true
		}
	}
}

// @lc code=end
