/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
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

func detectCycle(head *ListNode) *ListNode {
	var (
		fast *ListNode = head
		slow *ListNode = head
	)
	if head == nil {
		return nil
	}
	for {
		if fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return nil
		}
		if fast == slow {
			break
		}
	}
	for {
		if head == slow {
			return head
		}
		head, slow = head.Next, slow.Next
	}
}

// @lc code=end
