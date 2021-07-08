/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
package main

func multiply(num1 string, num2 string) string {
	buf := 0
	t := 0
	var r []int
	var rs []rune
	num1 = reverse(num1)
	num2 = reverse(num2)
	l1 := len(num1)
	l2 := len(num2)
	r = make([]int, l1+l2)
	rs = make([]rune, 0, l1+l2)
	for i := 0; i < l1; i++ {
		buf = 0
		for j := 0; j < l2; j++ {
			t = (int(num1[i])-'0')*(int(num2[j])-'0') + buf + r[i+j]
			r[i+j] = t % 10
			buf = t / 10
		}
		r[i+l2] = buf
	}
	// fmt.Print(r)
	for i, flag := l1+l2-1, false; i >= 0; i-- {
		if !flag && r[i] != 0 {
			flag = true
			rs = append(rs, rune(r[i])+'0')
		} else if flag {
			rs = append(rs, rune(r[i])+'0')
		}
	}
	// fmt.Print(rs)
	if len(rs) == 0 {
		return "0"
	}
	return string(rs)
}
func reverse(s string) string {
	runes := []rune(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

// func main() {
// 	fmt.Println(multiply("0", "0"))
// }

// @lc code=end
