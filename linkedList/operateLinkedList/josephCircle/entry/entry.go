package main

import (
	"fmt"
	"learnGo/linkedList/operateLinkedList/josephCircle"
)

func main() {
	list := josephCircle.List{}
	data := []josephCircle.Object{1,2,3,4,5,6,7,8}
	list.CreateLoopList(data)
	list.Traverse()

	fmt.Println("单向循环链表解决约瑟夫问题")
	list.
}
