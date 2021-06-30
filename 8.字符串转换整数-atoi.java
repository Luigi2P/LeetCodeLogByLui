/*
 * @lc app=leetcode.cn id=8 lang=java
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
class Solution {
    public int myAtoi(String s) {
        int t = 0;
        int state = 0;// 0前导空格 1符号或数字 2数字
        boolean minus = false;
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            switch (state) {
                case 0: {
                    if (c == ' ')
                        continue;
                    else {
                        state = 1;
                        i--;
                    }
                    break;
                }
                case 1: {
                    if (c == '-') {
                        minus = true;
                        state = 2;
                    } else if (c == '+') {
                        minus = false;
                        state = 2;
                    } else if (c >= '0' && c <= '9') {
                        state = 2;
                        t = c - '0';
                    } else
                        return t;
                    break;
                }
                case 2: {
                    if (c >= '0' && c <= '9') {
                        if (t > (Integer.MAX_VALUE - c + '0') / 10) {
                            return minus ? Integer.MIN_VALUE : Integer.MAX_VALUE;
                        }
                        t = t * 10 + c - '0';
                    } else
                        return t * (minus ? -1 : 1);
                }

            }
        }
        return t * (minus ? -1 : 1);
    }
}

// @lc code=end
// class Main {
// public static void main(String[] args) {
// Solution8 s = new Solution8();
// System.out.println(s.myAtoi("-91283472332"));
// }
// }