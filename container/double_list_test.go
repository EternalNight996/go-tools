// double_list_test.go
package container

import (
	"testing"
)

func TestDoubleList_Init(t *testing.T) {
	list := new(DoubleList)
	list.Init()
	if list.Length == 0 {
		t.Log("double list init success")
	} else {
		t.Error("double list init success")
	}
}

func TestDoubleList_Append(t *testing.T) {
	list := new(DoubleList)
	list.Init()

	flag := list.Append(nil)
	if flag {
		t.Error("double list append nil failed")
	} else {
		t.Log("double list append nil success")
	}

	flag = list.Append(&DoubleNode{DATA: 1})
	if flag && list.Length == 1 {
		t.Log("double list append 1 success")
	} else {
		t.Error("double list append 1 success")
	}

	list.Show()

	flag = list.Append(&DoubleNode{DATA: 2})
	if flag && list.Length == 2 {
		t.Log("double list append 2 success")
	} else {
		t.Error("double list append 2 success")
	}

	list.Show()
}

func TestDoubleList_Get(t *testing.T) {
	list := new(DoubleList)
	list.Init()

	node := list.Get(0)
	if node != nil {
		t.Error("empty double list get failed")
	}
	t.Log("empty double list get success")

	list.Append(&DoubleNode{DATA: 1})
	list.Append(&DoubleNode{DATA: 2})

	node = list.Get(0)
	if 1 == node.DATA.(int) {
		t.Log("double list get 1 success")
	} else {
		t.Error("double list get 1 error")
	}
	list.Show()

	node = list.Get(1)
	if 2 == node.DATA.(int) {
		t.Log("double list get 2 success")
	} else {
		t.Error("double list get 2 error")
	}
}

func TestDoubleList_Insert(t *testing.T) {
	list := new(DoubleList)
	list.Init()
	list.Show()
	list.ShowReverse()
	b := list.Insert(nil, 0)
	if b {
		t.Error("double list insert nil failed")
	} else {
		t.Log("double list insert nil success")
	}

	list.Insert(&DoubleNode{DATA: 1}, 0)
	node := list.Get(0)
	if 1 == node.DATA.(int) {
		t.Log("double insert get 1 success")
	} else {
		t.Error("double insert get 1 error")
	}

	list.Insert(&DoubleNode{DATA: 2}, 1)
	node = list.Get(1)
	if 2 == node.DATA.(int) {
		t.Log("double insert get 2 success")
	} else {
		t.Error("double insert get 2 error")
	}

	list.Insert(&DoubleNode{DATA: 3}, 1)
	node = list.Get(1)
	if 3 == node.DATA.(int) {
		t.Log("double insert get 3 success")
	} else {
		t.Error("double insert get 3 error")
	}

	list.Insert(&DoubleNode{DATA: 0}, 0)
	if list.Length == 4 {
		t.Log("double list insert header success")
	} else {
		t.Error("double list insert header failed")
	}

	list.Show()
	list.ShowReverse()
}

func TestDoubleList_Delete(t *testing.T) {
	list := new(DoubleList)
	list.Init()

	b := list.Delete(0)
	if b {
		t.Error("empty double list delete failed")
	} else {
		t.Log("empty double list delete success")
	}

	list.Append(&DoubleNode{DATA: 1})
	b = list.Delete(0)
	if b {
		t.Log("double list delete only one success")
	} else {
		t.Error("double list delete only one failed")
	}
	list.Append(&DoubleNode{DATA: 1})
	list.Append(&DoubleNode{DATA: 2})
	list.Append(&DoubleNode{DATA: 3})
	list.Append(&DoubleNode{DATA: 4})
	list.Append(&DoubleNode{DATA: 5})

	// delete head
	flag := list.Delete(0)
	if flag && list.Length == 4 {
		t.Log("double list delete 1 success")
	} else {
		t.Error("double list delete 1 error")
	}

	// delete tail
	flag = list.Delete(3)
	if flag && list.Length == 3 {
		t.Log("double list delete last success")
	} else {
		t.Error("double list delete last error")
	}

	// delete middle
	flag = list.Delete(1)
	if flag && list.Length == 2 {
		t.Log("double list delete middle success")
	} else {
		t.Error("double list delete middle error")
	}

	list.Show()
}
