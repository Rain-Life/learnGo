package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func variableFunc() {
	s := "你怕是个傻子"
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()

	count := utf8.RuneCountInString(s)
	fmt.Println(count)
	fmt.Println()

	bytes := []byte(s)
	for len(bytes) > 0 {
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

func arrayTest() {
	var hightBuilding [30]int
	for i:=0; i< 30; i++ {
		hightBuilding[i] = i+1
	}

	fmt.Println(hightBuilding[10:15])
	fmt.Println(hightBuilding[20:])
	fmt.Println(hightBuilding[2:])
}

func main()  {
	var a int
	var b float64
	var c string
	var d []int
	fmt.Println(a, b, c, d)

	test1 := 100
	test2 := 200
	test1, test2 = test2, test1
	fmt.Println(test1, test2)

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	const str = `first line
second line
third line
\r\n`
fmt.Println(str)

	variableFunc()
	arrayTest()
}
