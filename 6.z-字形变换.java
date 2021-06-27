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
        StringBuilder r = new StringBuilder();
        for (int i = 0; i < numRows; i++) {
            int j = 0;
            while (true) {
                if ((j * (2 * numRows - 2) + i) < leng)
                    r.append(s.charAt(j * (2 * numRows - 2) + i));
                else
                    break;
                if (i != 0 && i != numRows - 1 && ((j + 1) * (2 * numRows - 2) - i) < leng) {
                    r.append(s.charAt((j + 1) * (2 * numRows - 2) - i));
                }
                j++;
            }
        }
        return new String(r);
    }
}

// @lc code=end
class Main {
    public static void main(String[] args) {
        Solution6 solution6 = new Solution6();
        System.out.println(solution6.convert("PAYPALISHIRING", 1));
    }
}