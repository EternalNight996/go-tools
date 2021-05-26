//double_list.go
package container

import (
	"fmt"
	"sync"
)

// 节点数据
type DoubleFace interface{}

// 双链表节点
type DoubleNode struct {
	DATA DoubleFace
	PREV *DoubleNode
	NEXT *DoubleNode
}

// 双链表
type DoubleList struct {
	mutex  *sync.RWMutex
	Length uint
	HEAD   *DoubleNode
	TAIL   *DoubleNode
}

// 双链表初始化
func (list *DoubleList) Init() {
	//内存锁
	list.mutex = new(sync.RWMutex)
	list.Length = 0
	list.HEAD = nil
	list.TAIL = nil
}

// Get 获取指定位置的节点
func (list *DoubleList) Get(index uint) *DoubleNode {
	if list.Length == 0 || index > list.Length-1 {
		return nil
	} else if index == 0 {
		return list.HEAD
	}
	node := list.HEAD
	var i uint
	for i = 1; i <= index; i++ {
		node = node.NEXT
	}
	return node
}

// Append 向双链表后面追加节点
func (list *DoubleList) Append(node *DoubleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Length == 0 {
		list.HEAD = node
		list.TAIL = node
		node.NEXT = nil
		node.PREV = nil
	} else {
		node.PREV = list.TAIL
		node.NEXT = nil
		list.TAIL.NEXT = node
		list.TAIL = node
	}
	list.Length++
	return true
}

// Insert 向双链表指定位置插入节点
func (list *DoubleList) Insert(node *DoubleNode, index uint) bool {
	if index > list.Length || node == nil {
		return false
	}
	if index == list.Length {
		return list.Append(node)
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		node.NEXT = list.HEAD
		list.HEAD = node
		list.HEAD.PREV = nil
		list.Length++
		return true
	}

	nextNode := list.Get(index)
	node.PREV = nextNode.PREV
	node.NEXT = nextNode
	nextNode.PREV.NEXT = node
	nextNode.PREV = node
	list.Length++
	return true
}

// Delete 删除指定位置的节点
func (list *DoubleList) Delete(index uint) bool {
	if index+1 > list.Length {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		if list.Length == 1 {
			list.HEAD = nil
			list.TAIL = nil
		} else {
			list.HEAD.NEXT.PREV = nil
			list.HEAD = list.HEAD.NEXT
		}
		list.Length--
		return true
	}
	if index+1 == list.Length {
		list.TAIL.PREV.NEXT = nil
		list.TAIL = list.TAIL.PREV
		list.Length--
		return true
	}

	node := list.Get(index)
	node.PREV.NEXT = node.NEXT
	node.NEXT.PREV = node.PREV
	list.Length--
	return true
}

// Show打印双链表信息
func (list *DoubleList) Show() {
	if list == nil || list.Length == 0 {
		fmt.Println("这个双链表有错或为空！")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("这个双链表Flag>>%d\x0A", list.Length)
	ptr := list.HEAD
	for ptr != nil {
		fmt.Printf("数据>>%v\x0A", ptr.DATA)
		ptr = ptr.NEXT
	}
}

// ShowReverse 倒序打印双链表信息
func (list *DoubleList) ShowReverse() {
	if list == nil || list.Length == 0 {
		fmt.Println("这个双链表有错误或为空！")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("这个逆序双链表的Length>>%d\x0A", list.Length)
	ptr := list.TAIL
	for ptr != nil {
		fmt.Printf("数据>>%v\x0A", ptr.DATA)
		ptr = ptr.PREV
	}
}
