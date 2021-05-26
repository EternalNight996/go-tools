// stack.go
package container

// 栈信息
type Stack struct {
	list *SingleList
}

// Init 初始化栈
func (s *Stack) Init() {
	s.list = new(SingleList)
	s.list.Init()
}

// Push 压入栈
func (s *Stack) Push(data interface{}) bool {
	node := &SingleNode{
		DATA: data,
	}
	return s.list.Insert(node, 0)
}

// Pop 压出栈
func (s *Stack) Pop() interface{} {
	node := s.list.Get(0)
	if node != nil {
		s.list.Delete(0)
		return node.DATA
	}
	return nil
}

// Peek 查看栈顶元素
func (s *Stack) Peek() interface{} {
	node := s.list.Get(0)
	if node != nil {
		return node.DATA
	}
	return nil
}

// Size 获取栈的长度
func (s *Stack) Length() uint {
	return s.list.Length
}
