/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lA := getListLength(headA)
	lB := getListLength(headB)
	if lA > lB {
		lA, lB, headA, headB = lB, lA, headB, headA
	}
	diff := lB - lA
	for diff > 0 {
		headB = headB.Next
		diff--
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
func getListLength(head *ListNode) int {
	l := 0
	for head != nil {
		head = head.Next
		l++
	}
	return l
}

// @lc code=end
