/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	buf := 0
	t := 0
	r := "0"
	num1 = reverse(num1)
	num2 = reverse(num2)
	l1 := len(num1)
	l2 := len(num2)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			t = (int(num1[i])-'0')*(int(num2[j])-'0') + buf

		}
	}
}
func reverse(s string) string {
	runes := []rune(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

// @lc code=end
