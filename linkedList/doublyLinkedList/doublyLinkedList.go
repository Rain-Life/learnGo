package doublyLinkedList

//双向链表
type Object interface {}

type Node struct {
	Data Object
	Prev,Next *Node
}

type List struct {
	headNode *Node
}

//双向链表特点
/**
1、如果结点的Next指向为null，则说明是最后一个结点
*/
