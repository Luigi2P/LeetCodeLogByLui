import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
/*
 * @lc app=leetcode.cn id=1418 lang=java
 *
 * [1418] 点菜展示表
 */
import java.util.List;

// @lc code=start

class Solution {
    public List<List<String>> displayTable(List<List<String>> orders) {
        HashMap<String, HashMap<String, Integer>> H = new HashMap<>();// 表<桌名, 表<菜名, 数量>>
        HashSet<String> HSet = new HashSet<>();
        for (List<String> T : orders) {
            HashMap<String, Integer> HS = null;
            HSet.add(T.get(2));
            if (H.containsKey(T.get(1))) {
                HS = H.get(T.get(1));
                HS.put(T.get(2), HS.getOrDefault(T.get(2), 0) + 1);
            } else {
                HS = new HashMap<>();
                HS.put(T.get(2), 1);
                H.put(T.get(1), HS);
            }
        }
        ArrayList<String> foodItemList = new ArrayList<>(HSet);
        Collections.sort(foodItemList);
        List<List<String>> RList = new ArrayList<>();
        List<String> head = new ArrayList<>();
        head.add("Table");
        for (String i : foodItemList) {
            head.add(i);
        }
        RList.add(head);
        ArrayList<Integer> tableList = new ArrayList<>();
        for (String table : H.keySet()) {
            tableList.add(Integer.parseInt(table));
        }
        Collections.sort(tableList);
        for (int table : tableList) {
            ArrayList<String> T = new ArrayList<>();
            T.add(Integer.toString(table));
            for (String i : foodItemList) {
                if (H.get(Integer.toString(table)).get(i) != null)
                    T.add(H.get(Integer.toString(table)).get(i).toString());
                else
                    T.add("0");
            }
            RList.add(T);
        }
        return RList;
    }
}
// @lc code=end
