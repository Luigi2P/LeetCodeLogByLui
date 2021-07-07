/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package main

// import "fmt"

func search(nums []int, target int) int {
	return findR(nums, target, 0, len(nums)-1)
}

func findR(nums []int, target int, left int, right int) int {
	mid := (left + right) / 2
	if left == right {
		if nums[mid] == target {
			return mid
		} else {
			return -1
		}
	}
	if nums[left] <= nums[mid] { //左区间有序
		if target <= nums[mid] && target >= nums[left] { //target必在[left,mid]的条件
			return findR(nums, target, left, mid)
		} else {
			return findR(nums, target, mid+1, right)
		}
	} else { //右区间有序
		if target <= nums[right] && target >= nums[mid] { //target必在[mid,right]的条件
			return findR(nums, target, mid, right)
		} else {
			return findR(nums, target, left, mid-1)
		}
	}
}

// func main() {
// 	nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
// 	for i := 0; i < 8; i++ {
// 		fmt.Println(search(nums, i))
// 	}
// }

// @lc code=end
