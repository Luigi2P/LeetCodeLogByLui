/*
 * @lc app=leetcode.cn id=6 lang=java
 *
 * [6] Z 字形变换
 */

// @lc code=start
class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1)
            return s;
        int leng = s.length();
        int cyclelen = numRows * 2 - 2;
        StringBuilder r = new StringBuilder();
        for (int i = 0; i < numRows; i++) {
            for (int j = 0; j + i < leng; j += cyclelen) {
                r.append(s.charAt(j + i));
                if (i != 0 && i != numRows - 1 && j + cyclelen - i < leng)
                    r.append(s.charAt(j + cyclelen - i));
            }
        }
        return new String(r);
    }
}

// @lc code=end
// class Main {
// public static void main(String[] args) {
// Solution6 solution6 = new Solution6();
// System.out.println(solution6.convert("PAYPALISHIRING", 3));
// }
// }