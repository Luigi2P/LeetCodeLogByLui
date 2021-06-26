/*
 * @lc app=leetcode.cn id=2 lang=java
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list. public class ListNode { int val; ListNode
 * next; ListNode() {} ListNode(int val) { this.val = val; } ListNode(int val,
 * ListNode next) { this.val = val; this.next = next; } }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int t = 0;
        int T = 0;
        ListNode head = null;
        ListNode r = null;
        while (l1 != null || l2 != null || t != 0) {
            if (head == null) {
                r = new ListNode();
                head = r;
            } else {
                r.next = new ListNode();
                r = r.next;
            }
            T = t;
            if (l1 != null) {
                T += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                T += l2.val;
                l2 = l2.next;
            }
            r.val = T % 10;
            t = T / 10;
        }
        return head;
    }
}
// @lc code=end
