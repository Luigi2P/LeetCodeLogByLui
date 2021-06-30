/*
 * @lc app=leetcode.cn id=11 lang=java
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
class Solution {
    public int maxArea(int[] height) {
        int l = 0;
        int r = height.length - 1;
        int max = 0;
        int t;
        while (l != r) {
            t = Integer.min(height[l], height[r]) * (r - l);
            if (t > max)
                max = t;
            if (height[r] < height[l])
                r--;
            else
                l++;
        }
        return max;

    }
}
// @lc code=end
