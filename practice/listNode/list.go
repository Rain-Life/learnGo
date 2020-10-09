package main

import (
	"container/list"
	"fmt"
)

func insertValue() {
	l := list.New()

	l.PushBack("shulv")
	l.PushFront(88)

	element := l.PushBack("fist")

	l.InsertAfter("high", element)
	l.InsertBefore("noon", element)

	l.Remove(element)

	for i:=l.Front(); i!=nil; i=i.Next() {
		fmt.Println(i.Value)
	}
}


func main() {
	insertValue()
}
