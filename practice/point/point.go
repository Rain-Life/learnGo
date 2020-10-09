package main

import "fmt"

func swap(a,b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	a, b := 1, 2
	swap(&a,&b)
	fmt.Println(a, b)

	str := new(string)
	*str = "dhfghsdgfh"
	fmt.Println(*str)
}
