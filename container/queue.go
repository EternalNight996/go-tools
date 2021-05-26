// queue.go
package container

// Queue 队列信息
type Queue struct {
	list *SingleList
}

// Init 队列初始化
func (q *Queue) Init() {
	q.list = new(SingleList)
	q.list.Init()
}

// Length 获取队列长度
func (q *Queue) Length() uint {
	return q.list.Length
}

// Enqueue 进入队列
func (q *Queue) Enqueue(data interface{}) bool {
	return q.list.Append(&SingleNode{DATA: data})
}

// Dequeue 出列
func (q *Queue) Dequeue() interface{} {
	node := q.list.Get(0)
	if node == nil {
		return nil
	}
	q.list.Delete(0)
	return node.DATA
}

// Peek 查看队头信息
func (q *Queue) Peek() interface{} {
	node := q.list.Get(0)
	if node == nil {
		return nil
	}
	return node.DATA
}
