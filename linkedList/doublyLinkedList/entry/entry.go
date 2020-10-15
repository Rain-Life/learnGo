package main

import (
	"fmt"
	"learnGo/linkedList/doublyLinkedList"
)

func main() {
	list := doublyLinkedList.List{}

	fmt.Println("++++++++1、判断链表是否为空++++++++")
	fmt.Printf("链表是否为空：%v", list.IsEmpty())
	fmt.Println()
	fmt.Println()


}
