/*
 * @lc app=leetcode.cn id=7 lang=java
 *
 * [7] 整数反转
 */

// @lc code=start
class Solution {
    public int reverse(int x) {
        StringBuilder sb = new StringBuilder(String.valueOf(Math.abs(x)));
        if (x < 0)
            sb.append("-");
        sb.reverse();
        try {
            return Integer.parseInt(sb.toString());
        } catch (java.lang.NumberFormatException e) {
            return 0;
        }
    }
}

// @lc code=end
// class Main {
// public static void main(String[] args) {
// Solution7 s = new Solution7();
// System.out.println(s.reverse(1534236469));
// }
// }