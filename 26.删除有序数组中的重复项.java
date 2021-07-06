/*
 * @lc app=leetcode.cn id=26 lang=java
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
class Solution {
    public int removeDuplicates(int[] nums) {
        int l = 0;
        int r = 1;
        int length = nums.length;
        if (length < 2)
            return length;
        while (r < length) {
            if (nums[l] != nums[r]) {
                nums[l + 1] = nums[r];
                l++;
            }
            r++;
        }
        return l + 1;
    }
}
// @lc code=end
