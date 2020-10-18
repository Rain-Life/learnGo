package main

import (
	"fmt"
	"learnGo/linkedList/doublyLoopLinkedList"
)

func print(list doublyLoopLinkedList.List)  {
	fmt.Printf("链表长度%d\n", list.Length())
	fmt.Println("遍历链表所有结点：")
	list.Traverse()
	fmt.Println()
	fmt.Println()
}

func main() {
	list := doublyLoopLinkedList.List{}

	fmt.Println("++++++++1、判断链表是否为空++++++++")
	fmt.Printf("链表是否为空：%v", list.IsEmpty())
	fmt.Println()
	fmt.Println()

	fmt.Println("++++++++2、向双向链表头部添加元素++++++++")
	fmt.Println()
	list.AddFromHead(1)
	list.AddFromHead(2)
	list.AddFromHead(3)
	fmt.Println("44444")
	print(list)
}
