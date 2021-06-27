/*
 * @lc app=leetcode.cn id=5 lang=java
 *
 * [5] 最长回文子串
 */

// @lc code=start
class Solution {
    public String longestPalindrome(String s) {
        // 动态规划
        int length = s.length();
        boolean[][] dp = new boolean[length][length];
        int begin = 0;
        int maxlength = 1;
        if (length < 2)
            return s;
        for (int i = 0; i < length; i++)
            dp[i][i] = true;
        char[] sArray = s.toCharArray();
        for (int L = 2; L <= length; L++) {
            for (int i = 0; i < length; i++) {
                int j = L + i - 1;
                if (j >= length)
                    break;
                if (sArray[i] == sArray[j] && (dp[i + 1][j - 1] || i + 1 == j)) {
                    dp[i][j] = true;
                    if ((j - i + 1) > maxlength) {
                        maxlength = j - i + 1;
                        begin = i;
                    }
                }
            }
        }

        return s.substring(begin, begin + maxlength);
    }
}
// @lc code=end

// class Main {
// public static void main(String[] args) {
// Solution5 s = new Solution5();
// System.out.print(s.longestPalindrome("cbdddddbc"));
// }
// }