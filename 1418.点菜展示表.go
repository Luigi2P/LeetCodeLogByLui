/*
 * @lc app=leetcode.cn id=1418 lang=golang
 *
 * [1418] 点菜展示表
 */

// @lc code=start
package main

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string { //[1]桌名 [2]食物名
	H := map[string](map[string]int){} //表<桌名, 表<菜名, 数量>>
	HSet := map[string]struct{}{}      //食物Set
	for _, T := range orders {
		HSet[T[2]] = struct{}{}
		_, ok := H[T[1]]
		if ok {
			H[T[1]][T[2]]++
		} else {
			H[T[1]] = map[string]int{T[2]: 1}
		}
	}
	var foodItemList []string
	for key := range HSet {
		foodItemList = append(foodItemList, key)
	}
	sort.Strings(foodItemList)
	var RList [][]string
	var head []string = []string{"Table"}
	head = append(head, foodItemList...)
	RList = append(RList, head)
	var tableList []int
	for key := range H {
		n, _ := strconv.Atoi(key)
		tableList = append(tableList, n)
	}
	sort.Ints(tableList)
	for _, table := range tableList {
		var T []string
		T = append(T, strconv.Itoa(table))
		for _, value := range foodItemList {
			v, ok := H[strconv.Itoa(table)][value]
			if ok {
				T = append(T, strconv.Itoa(v))
			} else {
				T = append(T, "0")
			}
		}
		RList = append(RList, T)
	}

	return RList
}

// @lc code=end
// func main() {
// 	orders := [][]string{{"111", "2", "Beef Burrito"}, {"Jhon", "1", "Aeer Burrito"}, {"Melissa", "2", "Soda"}}
// 	displayTable(orders)
// }
