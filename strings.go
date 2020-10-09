package main

import (
	"fmt"
	"unicode/utf8"
)

func stringByte(s string) {
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)//每个中文（UTF-8编码）占三个字节
	}
	fmt.Println()

	for i, ch := range s { // ch 是 rune类型的
		fmt.Printf("(%d, %X) ", i, ch) //上边打印的b是UTF-8的编码，这里ch打印出来的是Unicode编码
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s))//11

	bytes := []byte(s)
	for len(bytes) >0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}
	fmt.Println()
}

func main() {
	s := "开始学习Golang!"
	stringByte(s)
}
