package main

import (
	"fmt"
	"learnGo/retriver/mock"
	realR "learnGo/retriver/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake www.baidu.com"}
	fmt.Printf("%T, %v", r, r)

	fmt.Println()

	r = realR.Retriever{}
	fmt.Printf("%T, %v", r, r)
	//fmt.Println(download(r))
}
