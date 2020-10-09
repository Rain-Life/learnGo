package main

import "fmt"

func learnSlice() {
	arr := [...]int{0,1,2,3,4,5,6,7,8}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:6])
	fmt.Println("arr[:] = ", arr[:])
}

func updateSlice(s []int) {
	s[0] = 520
}

func learnSlice1() {
	arr := [...]int{0,1,2,3,4,5,6,7,8}
	s1 := arr[2:6]
	s2 := s1[1:]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
}

func learnSlice2() {
	arr := [...]int{0,1,2,3,4,5,6,7,8}
	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))
	fmt.Println()
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}


func main() {
	learnSlice1()

	arr := [...]int{0,1,2,3,4,5,6,7,8}
	s1 := arr[2:]
	fmt.Println("s1更新之前")
	fmt.Println("s1 = ", s1)
	fmt.Println("arr = ", arr)

	fmt.Println("更新切片s1")
	updateSlice(s1)

	fmt.Println("更新之后，s1和arr的值")
	fmt.Println(s1)
	fmt.Println(arr)

	learnSlice2()
}
