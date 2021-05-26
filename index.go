// index.go
package main

import (
	"fmt"

	"github.com/EternalNight996/go-tools/container"
)

func main() {
	fmt.Println("\x0A----------测试单链表入栈stack----------")
	stack := new(container.Stack)
	stack.Init()
	stack.Push("queue-1")
	stack.Push("queue-2")
	stack.Push("queue-3")
	fmt.Println("Length: ", stack.Length())
	fmt.Println("Peek: ", stack.Peek()) //把头队列queue-1找出来
	fmt.Println("Pop: ", stack.Pop())   //出队列从queue-1开始出
	fmt.Println("Length: ", stack.Length())

	fmt.Println("\x0A----------测试单链表队列queue----------")
	queue := new(container.Queue)
	queue.Init()
	queue.Enqueue("queue-1") //入队列
	queue.Enqueue("queue-2")
	queue.Enqueue("queue-3")
	fmt.Println("Length: ", queue.Length())
	fmt.Println("Peek: ", queue.Peek())       //把头队列queue-1找出来
	fmt.Println("Dequeue: ", queue.Dequeue()) //出队列从queue-1开始出
	fmt.Println("Length: ", queue.Length())

	fmt.Println("\x0A----------测试单链表single_list----------")
	single_list := new(container.SingleList)
	single_list.Init()
	single_list.Append(&container.SingleNode{DATA: "SingleAppend-1"})
	single_list.Append(&container.SingleNode{DATA: "SingleAppend-2"})
	single_list.Append(&container.SingleNode{DATA: "SingleAppend-3"})
	single_list.Show() //从头部开始提取数据 从SingleNode>SingleFace.....
	single_list.Insert(&container.SingleNode{DATA: "消灭人类暴政，世界属于三体."}, 1)
	fmt.Println("Get[1]: ", single_list.Get(1).DATA) //DATA是指向SingleFace的
	fmt.Println("Length是记录总长度： ", single_list.Length)
	single_list.Delete(1) //删除[1]
	single_list.Show()

	fmt.Println("\x0A----------测试双链表double_list----------")
	double_list := new(container.DoubleList)
	double_list.Init()
	double_list.Append(&container.DoubleNode{DATA: "DoubleAppend-1"})
	double_list.Append(&container.DoubleNode{DATA: "DoubleAppend-2"})
	double_list.Append(&container.DoubleNode{DATA: "DoubleAppend-3"})
	double_list.Append(&container.DoubleNode{DATA: "DoubleAppend-4"})
	double_list.Show()
	double_list.Insert(&container.DoubleNode{DATA: "消灭人类暴政，世界属于三体."}, 4)
	fmt.Println("Get[1]: ", double_list.Get(1).DATA)
	fmt.Println("Length是记录总长度： ", double_list.Length)
	double_list.Delete(4)
	double_list.ShowReverse() //双链表多了个逆序,差异在那？DoubleNode多了PREV指向DoubleFace相当上层
}
