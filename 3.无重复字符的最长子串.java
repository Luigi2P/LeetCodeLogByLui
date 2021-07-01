/*
 * @lc app=leetcode.cn id=3 lang=java
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
// import java.util.HashMap;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int l = s.length();
        if (l == 0)
            return 0;
        int n = 0;
        int maxn = 0;
        HashMap<Character, Integer> map = new HashMap<Character, Integer>();
        for (int i = 0; i < l; i++) {
            if (map.containsKey(s.charAt(i))) {
                int t = i - map.get(s.charAt(i));
                n = t <= n ? t : n + 1;
                map.replace(s.charAt(i), i);
            } else {
                n++;
                map.put(s.charAt(i), i);
            }
            maxn = n > maxn ? n : maxn;
        }
        return maxn;
    }
}

// import java.util.HashSet;
// class Solution {
// public int lengthOfLongestSubstring(String s) {
// // 哈希集合，记录每个字符是否出现过
// Set<Character> occ = new HashSet<Character>();
// int n = s.length();
// // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
// int rk = -1, ans = 0;
// for (int i = 0; i < n; ++i) {
// if (i != 0) {
// // 左指针向右移动一格，移除一个字符
// occ.remove(s.charAt(i - 1));
// }
// while (rk + 1 < n && !occ.contains(s.charAt(rk + 1))) {
// // 不断地移动右指针
// occ.add(s.charAt(rk + 1));
// ++rk;
// }
// // 第 i 到 rk 个字符是一个极长的无重复字符子串
// ans = Math.max(ans, rk - i + 1);
// }
// return ans;
// }
// }
// @lc code=end
// class Main {
// public static void main(String[] args) {
// Solution3 s = new Solution3();
// s.lengthOfLongestSubstring("tmmzuxt");
// }
// }