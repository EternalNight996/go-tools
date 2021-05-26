// single_list_test.go
package container

import (
	"testing"
)

func TestSingleList_Init(t *testing.T) {
	list := new(SingleList)
	list.Init()
	t.Log("single list init success")
}

func TestSingleList_Append(t *testing.T) {
	list := new(SingleList)
	list.Init()

	b := list.Append(nil)
	if b {
		t.Error("single list append nil failed")
	} else {
		t.Log("single list append nil success")
	}

	b = list.Append(&SingleNode{DATA: 1})
	if b {
		t.Log("single list append first success")
	} else {
		t.Error("single list append first failed")
	}

	b = list.Append(&SingleNode{DATA: 2})
	if b {
		t.Log("single list append second success")
	} else {
		t.Error("single list append second failed")
	}
}

func TestSingleList_Insert(t *testing.T) {
	list := new(SingleList)
	list.Init()

	b := list.Insert(nil, 0)
	if b {
		t.Error("single list insert nil failed")
	} else {
		t.Log("single list insert nil success")
	}

	b = list.Insert(&SingleNode{DATA: 1}, 1)
	if b {
		t.Error("single list insert out of range failed")
	} else {
		t.Log("single list insert out of range success")
	}

	b = list.Insert(&SingleNode{DATA: 1}, 0)
	if b {
		t.Log("single list insert first success")
	} else {
		t.Error("single list insert first failed")
	}

	b = list.Insert(&SingleNode{DATA: 2}, 1)
	b = list.Insert(&SingleNode{DATA: 3}, 2)
	if b {
		t.Log("single list insert multi success")
	} else {
		t.Error("single list insert multi failed")
	}
}

func TestSingleList_Delete(t *testing.T) {
	list := new(SingleList)
	list.Init()

	b := list.Delete(0)
	if b {
		t.Error("single list delete out of range failed")
	} else {
		t.Log("single list delete out of range success")
	}

	list.Append(&SingleNode{DATA: 1})

	b = list.Delete(0)
	if b {
		t.Log("single list delete first success")
	} else {
		t.Error("single list delete first failed")
	}

	list.Append(&SingleNode{DATA: 1})
	list.Append(&SingleNode{DATA: 2})
	list.Append(&SingleNode{DATA: 3})

	b = list.Delete(2)
	if b {
		t.Log("single list delete third success")
	} else {
		t.Error("single list delete third failed")
	}
}

func TestSingleList_Get(t *testing.T) {
	list := new(SingleList)
	list.Init()

	node := list.Get(0)
	if node == nil {
		t.Log("empty single list get success")
	} else {
		t.Error("empty single list get failed")
	}

	list.Append(&SingleNode{DATA: 1})
	list.Append(&SingleNode{DATA: 2})

	node = list.Get(0)
	if 1 == node.DATA.(int) {
		t.Log("single list get first success")
	} else {
		t.Error("single list get first failed")
	}

	node = list.Get(1)
	if 2 == node.DATA.(int) {
		t.Log("single list get second success")
	} else {
		t.Error("single list get second failed")
	}
}

func TestSingleList_Show(t *testing.T) {
	list := new(SingleList)
	list.Init()

	list.Show()

	list.Append(&SingleNode{DATA: 1})
	list.Append(&SingleNode{DATA: 2})
	list.Append(&SingleNode{DATA: 3})
	list.Show()
}
