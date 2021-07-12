/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */
package main

// @lc code=start
func reverseWords(s string) string {
	var l, r int
	l = -1
	var result = make([]byte, 0, len(s))
	for r != len(s) {
		if s[r] == ' ' {
			for i := r - 1; i > l; i-- {
				result = append(result, byte(s[i]))
			}
			result = append(result, ' ')
			l = r
		}
		r++
	}
	for i := r - 1; i > l; i-- {
		result = append(result, byte(s[i]))
	}
	return string(result)
}

// @lc code=end
