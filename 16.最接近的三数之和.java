// import java.util.Arrays;
/*
 * @lc app=leetcode.cn id=16 lang=java
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        int length = nums.length;
        Arrays.sort(nums);
        int result = 0;
        int min = Integer.MAX_VALUE;
        for (int i = 0; i < length; i++) {
            if (i == 0 || nums[i - 1] != nums[i]) {
                int l = i + 1;
                int r = length - 1;
                while (l < r) {
                    int sum = nums[i] + nums[l] + nums[r];
                    if (sum == target)
                        return sum;
                    if (Math.abs(sum - target) < min) {
                        min = Math.abs(sum - target);
                        result = sum;
                    }
                    if (sum > target)
                        r--;
                    else
                        l++;
                }
            }
        }
        return result;
    }
}

// @lc code=end
// class Main {
// public static void main(String[] args) {
// Solution16 s = new Solution16();
// int[] I = { 1, 1, 1, 0 };
// System.out.println(s.threeSumClosest(I, -100));
// }
// }