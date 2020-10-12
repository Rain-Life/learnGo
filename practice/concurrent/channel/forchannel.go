package channel

import (
	"fmt"
	"time"
)

func Forchannel() {
	ch := make(chan int)
	go OtherGoRoutine(ch)

	fmt.Println("主函数goroutine在执行...")
	for data := range ch {
		fmt.Println("打印channel中的值：", data)

		if data == 0 {
			break
		}
	}
	fmt.Println("全部执行结束...")
}


func OtherGoRoutine(ch chan int) {
	fmt.Println("匿名函数goroutine被执行...")
	for i:=3; i>=0; i-- {
		fmt.Println("匿名函数循环执行", i)
		ch <- i
		time.Sleep(time.Second)
	}

	fmt.Println("匿名函数goroutine执行结束...")
}
