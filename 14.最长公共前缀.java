/*
 * @lc app=leetcode.cn id=14 lang=java
 *
 * [14] 最长公共前缀
 */

// @lc code=start
class Solution {
    public String longestCommonPrefix(String[] strs) {
        int l = strs.length;
        if (l < 2)
            return strs[0];
        String s = strs[0];
        for (int i = 1; i < l; i++) {
            s = longestCommonPrefixBetweenTwo(s, strs[i]);
            if (s.length() == 0)
                return s;
        }
        return s;
    }

    public String longestCommonPrefixBetweenTwo(String a, String b) {
        StringBuilder r = new StringBuilder();
        int l = Math.min(a.length(), b.length());
        if (l == 0)
            return "";
        for (int i = 0; i < l; i++) {
            if (a.charAt(i) == b.charAt(i))
                r.append(a.charAt(i));
            else
                return r.toString();
        }
        return r.toString();
    }
}
// @lc code=end
