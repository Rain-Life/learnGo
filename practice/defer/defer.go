package main

import (
	"fmt"
	"os"
	"sync"
)

func practiceDefer1() {
	fmt.Println("start defer...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end defer...")
}

var (
	valueByKey = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

func readValue1(key string) int{
	valueByKeyGuard.Lock()
	v:=valueByKey[key]
	valueByKeyGuard.Unlock()
	return v
}

func readValue2(key string) int  {
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func fileSize1(filename string) int64  {
	f, err := os.Open(filename)
	if err!=nil {
		return 0
	}

	info, err := f.Stat()
	if err != nil {
		f.Close()
		return 0
	}
	size := info.Size()
	f.Close()

	return size
}

func fileSize2(filename string) int64 {
	f, err:=os.Open(filename)
	if err!=nil {
		return 0
	}

	defer f.Close()
	info, err:=f.Stat()
	if err!=nil {
		return 0
	}

	size := info.Size()

	return size
}

func main() {
	practiceDefer1()
}
