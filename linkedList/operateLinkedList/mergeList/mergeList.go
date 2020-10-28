package mergeList

import "fmt"

type Object interface {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

//生成单链表
func (list *List) CreateList(data []Object) {
	count := len(data)

	if count < 0 || count > 1000 {
		fmt.Println("参数不合法")
		return
	} else if count == 1 {
		node := &Node{Data: data[0]}
		list.headNode = node
	} else {
		for _, value := range data {
			node := &Node{Data: value}
			//为了方便，这里使用头插法
			node.Next = list.headNode
			list.headNode = node
		}
	}
}

//遍历链表
func (list List) Traverse() {
	if list.headNode == nil {
		fmt.Println("链表为空")
		return
	}

	currentNode := list.headNode
	for currentNode != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
}




