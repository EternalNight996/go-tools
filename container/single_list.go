// single_list.go
package container

import (
	"fmt"
	"sync"
)

// 节点数据
type SingleFace interface{}

// 单链表节点
type SingleNode struct {
	DATA SingleFace
	NEXT *SingleNode
}

// 单链表
type SingleList struct {
	// 互斥锁==读写锁
	mutex  *sync.RWMutex
	HEAD   *SingleNode
	TAIL   *SingleNode
	Length uint
}

// 初始化
func (list *SingleList) Init() {
	list.Length = 0
	list.HEAD = nil
	list.TAIL = nil
	list.mutex = new(sync.RWMutex)
}

// 添加节点到链表尾部
func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Length == 0 {
		list.HEAD = node
		list.TAIL = node
		list.Length = 1
		return true
	}

	list.TAIL.NEXT = node
	list.TAIL = node
	list.Length += 1
	return true
}

// 插入节点到指定位置
func (list *SingleList) Insert(node *SingleNode, index uint) bool {
	if node == nil || index > list.Length {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		node.NEXT = list.HEAD
		list.HEAD = node
		list.Length += 1
		return true
	}
	var i uint
	ptr := list.HEAD
	for i = 1; i < index; i++ {
		ptr = ptr.NEXT
	}
	next := ptr.NEXT
	ptr.NEXT = node
	node.NEXT = next
	list.Length += 1
	return true
}

// 删除指定位置的节点
func (list *SingleList) Delete(index uint) bool {
	if list == nil || list.Length == 0 || index > list.Length-1 {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		head := list.HEAD.NEXT
		list.HEAD = head
		if list.Length == 1 {
			list.TAIL = nil
		}
		list.Length -= 1
		return true
	}

	ptr := list.HEAD
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.NEXT
	}
	next := ptr.NEXT

	ptr.NEXT = next.NEXT
	if index == list.Length-1 {
		list.TAIL = ptr
	}
	list.Length -= 1
	return true
}

// 获取指定位置的节点，不存在则返回nil
func (list *SingleList) Get(index uint) *SingleNode {
	if list == nil || list.Length == 0 || index > list.Length-1 {
		return nil
	}

	list.mutex.RLock()
	defer list.mutex.RUnlock()

	if index == 0 {
		return list.HEAD
	}
	node := list.HEAD
	var i uint
	for i = 0; i < index; i++ {
		node = node.NEXT
	}
	return node
}

// 输出链表
func (list *SingleList) Show() {
	if list == nil || list.Length == 0 {
		fmt.Println("这个单链表为空!")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("这个单链表长度为%d\x0A", list.Length)
	ptr := list.HEAD
	var i uint
	for i = 0; i < list.Length; i++ {
		fmt.Printf("FLAG:%3d 数据为%v\x0A", i+1, ptr.DATA)
		ptr = ptr.NEXT
	}
}
