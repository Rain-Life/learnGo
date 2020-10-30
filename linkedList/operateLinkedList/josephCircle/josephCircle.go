package josephCircle

import "fmt"

type Object interface {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

//遍历链表
func (list List)Traverse()  {
	if list.headNode == nil {
		fmt.Printf("链表为空")
		return
	}

	currentNode := list.headNode
	for currentNode.Next != list.headNode {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Printf("%v\t", currentNode.Data)
}

//生成循环链表
func (list *List)CreateLoopList(data []Object)  {
	if len(data) < 0 {
		fmt.Println("参数不合法")
		return
	}

	currentNode := list.headNode
	for i, value := range data {
		node := &Node{Data: value}
		if i == 0 {
			list.headNode = node
			currentNode = node
			continue
		}
		currentNode.Next = node
		currentNode = node
	}
	currentNode.Next = list.headNode
}

//单向链表解决约瑟夫问题
func (list *List) dealJosephCircle(number int) []interface{} {
	//需要考虑链表只有一个节点的情况
	var data []interface{}
	index := 1
	currentNode := list.headNode
	//preNode := list.headNode

	for list.headNode != nil {
		//只有一个节点的情况
		if currentNode.Next == list.headNode {
			data = append(data, currentNode.Data)
			list.headNode = nil
			continue
		}

		if index == number {
			//删除结点，其实不用考虑是不是头结点或尾结点


			index = 0
		} else {
			index++
		}

		currentNode = currentNode.Next
	}
	return data
}
