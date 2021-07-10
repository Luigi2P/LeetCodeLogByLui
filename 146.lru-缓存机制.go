/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */
package main

// @lc code=start
type LRUCache struct {
	capacity, size int
	head, tail     *DLinkedNode
	cache          map[int]*DLinkedNode
}
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
		cache:    map[int]*DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.next = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if ok {
		this.moveToHead(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.cache[key]
	if ok {
		v.value = value
		this.moveToHead(v)
		return
	}
	if this.capacity == this.size {
		t := this.tail.prev
		this.removeTail()
		delete(this.cache, t.key)
		this.size--
	}
	l := &DLinkedNode{
		key:   key,
		value: value,
	}
	this.addToHead(l)
	this.size++
	this.cache[key] = l
}
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
func (this *LRUCache) remove(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}
func (this *LRUCache) removeTail() {
	this.remove(this.tail.prev)
}
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.remove(node)
	this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Get(3)    // 返回 3
	lRUCache.Get(4)    // 返回 4
}
