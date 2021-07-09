/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
package main

func generateMatrix(n int) [][]int {

	var (
		r              = make([][]int, n)
		uB, dB, lB, rB = 0, n - 1, 0, n - 1
		x, y           = 0, -1
		index          = 0
	)
	for i := range r {
		r[i] = make([]int, n)
	}
	for {
		for y < rB {
			y++
			index++
			r[x][y] = index
		}
		uB++
		if uB > dB {
			break
		}
		for x < dB {
			x++
			index++
			r[x][y] = index
		}
		rB--
		if rB < lB {
			break
		}
		for y > lB {
			y--
			index++
			r[x][y] = index
		}
		dB--
		if dB < uB {
			break
		}
		for x > uB {
			x--
			index++
			r[x][y] = index
		}
		lB++
		if lB > rB {
			break
		}

	}
	return r
}

// @lc code=end
