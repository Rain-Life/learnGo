package channel

import "fmt"

func Test() {
	ch := make(chan int, 3)
	fmt.Println("当前通道大小：", len(ch))

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(len(ch))
}
