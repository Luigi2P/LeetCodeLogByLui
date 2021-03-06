package main

import "fmt"

func main() {
	nums := []int{7, 5, 3, 6, 2, 4, 22, 6, 89}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	t := nums[left]
	for i < j {
		for i < j && nums[j] <= t {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] >= t {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = t
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
