/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
package main

func spiralOrder(matrix [][]int) []int {
	var (
		m              = len(matrix[0])
		n              = len(matrix)
		r              = make([]int, 0, m*n)
		uB, dB, lB, rB = 0, n - 1, 0, m - 1
		x, y           = 0, -1
	)
	for {
		for y < rB {
			y++
			r = append(r, matrix[x][y])
		}
		uB++
		if uB > dB {
			break
		}
		for x < dB {
			x++
			r = append(r, matrix[x][y])
		}
		rB--
		if rB < lB {
			break
		}
		for y > lB {
			y--
			r = append(r, matrix[x][y])
		}
		dB--
		if dB < uB {
			break
		}
		for x > uB {
			x--
			r = append(r, matrix[x][y])
		}
		lB++
		if lB > rB {
			break
		}

	}
	return r
}

// func main() {
// 	// t := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
// 	// t := [][]int{{1}}
// 	t := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	fmt.Println(spiralOrder((t)))
// }

// @lc code=end
