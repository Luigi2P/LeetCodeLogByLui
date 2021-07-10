/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	minS  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	l := MinStack{
		stack: []int{},
		minS:  []int{math.MaxInt64},
	}
	return l
}

func (this *MinStack) Push(val int) {
	if val <= this.GetMin() {
		this.minS = append(this.minS, val)
	} else {
		this.minS = append(this.minS, this.GetMin())
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.minS = this.minS[:len(this.minS)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minS[len(this.minS)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
func main() {
	ms := Constructor()
	ms.Push(0)
	ms.Push(1)
	ms.Push(0)
	fmt.Println(ms.GetMin())
}
