
/*
 * @lc app=leetcode.cn id=15 lang=java
 *
 * [15] 三数之和
 */
// import java.util.List;
// import java.util.ArrayList;
// import java.util.Arrays;

// @lc code=start
class Solution15 {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> L = new ArrayList<>();
        int l = nums.length;
        Arrays.sort(nums);
        if (l < 3)
            return L;
        int k;
        for (int i = 0; i < l; i++) {
            k = l - 1;
            if (i == 0 || nums[i - 1] != nums[i])
                for (int j = i + 1; j < l; j++) {
                    if (j == i + 1 || nums[j - 1] != nums[j]) {
                        int t = -nums[i] - nums[j];
                        while (t < nums[k] && k > 0)
                            k--;
                        if (k <= j)
                            break;
                        if (nums[k] == t) {
                            List<Integer> TL = new ArrayList<>();
                            TL.add(nums[i]);
                            TL.add(nums[j]);
                            TL.add(t);
                            L.add(TL);
                        }
                    }
                }
        }
        return L;
    }
}

// @lc code=end
// class Main {
// public static void main(String[] args) {
// Solution15 s = new Solution15();
// int[] nums = { -1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4 };
// System.out.println(s.threeSum(nums));
// }
// }