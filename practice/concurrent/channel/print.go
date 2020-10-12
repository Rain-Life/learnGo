package channel

import "fmt"

func printer(c chan int) {
	for  {
		data := <-c
		if data == 0 {
			break
		}
		fmt.Println("printer函数打印通道值", data)
	}
	//fmt.Println("我是printer函数，我执行结束了...")
	//c <- 0
}

func ExecFunc() {
	c := make(chan int)

	go printer(c)

	for i:=1; i<=10; i++ {
		c <- i
	}

	fmt.Println("告诉printer，没有数据了")
	c <- 0
	//fmt.Println("printer，你要是执行完了告诉我")
	//<-c
}
