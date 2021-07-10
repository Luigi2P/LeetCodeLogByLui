/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// tail := head
	// if head == nil {
	// 	return nil
	// }
	// for tail.Next != nil {
	// 	tail = tail.Next
	// }
	// for head != tail {
	// 	t := head
	// 	head = head.Next
	// 	t.Next = tail.Next
	// 	tail.Next = t
	// }
	// return tail
	if head == nil {
		return nil
	}
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// @lc code=end

