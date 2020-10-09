package main

import "fmt"

func testFor1() {
	for i:=1; i<10; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, i*j)
			if j==8 {
				goto breakHere
			}
		}
		fmt.Println()
	}

	breakHere:
		fmt.Println("Now j=8")
}

func testSwitch() {
	var s = "hello"
	switch s {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("this is default")
	}
}

func main() {
	testFor1()
	testSwitch()
}
