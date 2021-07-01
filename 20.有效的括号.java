import java.util.Stack;

/*
 * @lc app=leetcode.cn id=20 lang=java
 *
 * [20] 有效的括号
 */

// @lc code=start
class Solution {
    public boolean isValid(String s) {
        // 0() 1{} 2[]
        Stack<Integer> stack = new Stack<>();
        int length = s.length();
        for (int i = 0; i < length; i++) {
            switch (s.charAt(i)) {
                case '(': {
                    stack.push(0);
                    break;
                }
                case ')':
                    if (stack.empty() || stack.pop() != 0)
                        return false;
                    break;
                case '[': {
                    stack.push(1);
                    break;
                }
                case ']':
                    if (stack.empty() || stack.pop() != 1)
                        return false;
                    break;
                case '{': {
                    stack.push(2);
                    break;
                }
                case '}':
                    if (stack.empty() || stack.pop() != 2)
                        return false;
                    break;
                default:
                    return false;
            }
        }
        return stack.empty();
    }
}
// @lc code=end
