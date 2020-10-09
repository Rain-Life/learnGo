package main

import "fmt"

func visit(list []int, f func(int)) {
	for _,v := range list {
		f(v)
	}
}

func Accumulate(value int) func() int  {
	return func() int{
		value++
		return value
	}
}

func main() {
	f:=func(data int) {
		fmt.Println("hello ", data)
	}
	f(100)
	visit([]int{1,2,3}, func(v int) {
		fmt.Println(v)
	})
	str := "hello world"
	foo := func() {
		str = "hello shulv"
	}
	foo()
	fmt.Println(str)

	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)
}
