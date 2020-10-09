package main

import (
	"fmt"
	"sort"
	"sync"
	"unicode/utf8"
)

func mapTest1() {
	sence := make(map[string]int)
	sence["route"] = 100
	fmt.Println(sence["route"])

	v := sence["route2"]
	fmt.Println(v)
}

func mapTest2()  {
	sence := map[string]int{
		"test1" : 10,
		"test2" : 20,
		"test3" : 30,
	}
	for k, v := range sence {
		fmt.Println(k, v)
	}
}

func mapTest3() {
	sence := make(map[string]int)
	sence["route"] = 66
	sence["brazil"] = 4
	sence["china"] = 960

	var senceList []string
	for k := range sence {
		senceList = append(senceList, k)
	}

	sort.Strings(senceList)
	fmt.Println(senceList)

	delete(sence, "brazil")
	for k,v:=range sence {
		fmt.Println(k, v)
	}
}

func mapTest4() {
	var sence sync.Map
	sence.Store("greece", 97)
	sence.Store("london", 100)
	sence.Store("egypt", 200)

	fmt.Println(sence.Load("london"))

	sence.Range(func(k, v interface{}) bool {
		fmt.Println("iterate: ", k, v)
		return true
	})
}

//寻找最长不含有重复字符的子串
func noneRepeating(s string) int {
	lastOccurred := make(map[byte]int)
	start:=0
	maxLength:=0
	for i,ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI+1
		}
		if i-start+1 > maxLength {
			maxLength = i-start+1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func runeType() {
	s:="IT猿圈!"
	fmt.Println(len(s))
	fmt.Printf("%s\n", s)
	for _,b := range []byte(s) { //UTF-8
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i,ch := range s {// ch is a rune (Unicode)
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
}

func newNoneRepeating(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength:=0
	for i, ch:=range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI>=start {
			start = lastI+1
		}
		if i-start+1 > maxLength {
			maxLength = i-start+1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	//mapTest1()
	//mapTest2()
	//mapTest3()
	//mapTest4()
	//fmt.Println(noneRepeating("abcabcbb"))
	//fmt.Println(noneRepeating("bbbbb"))
	//fmt.Println(noneRepeating("pwwkew"))
	//runeType()
	fmt.Println(newNoneRepeating("我爱喝芬达"))
}
