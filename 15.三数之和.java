
/*
 * @lc app=leetcode.cn id=15 lang=java
 *
 * [15] 三数之和
 */
import java.util.List;
import java.util.ArrayList;

// @lc code=start
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> L = new ArrayList<>();
        int l = nums.length;
        if (l < 3)
            return L;
        for (int i = 0; i < l; i++) {
            for (int j = i + 1; j < l; j++) {
                int t = -nums[i] - nums[j];
                for (int k = j + 1; k < l; k++) {
                    if (nums[k] == t) {
                        List<Integer> TAL = new ArrayList<>();
                        TAL.add(nums[i]);
                        TAL.add(nums[j]);
                        TAL.add(nums[k]);
                        L.add(TAL);
                    }
                }
            }
        }
        return L;

    }
}
// @lc code=end
