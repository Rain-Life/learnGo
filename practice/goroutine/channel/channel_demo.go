package goroutine

import "fmt"

func CreateWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			n:=<-c
			fmt.Printf("Worker %d received %d\n",id, n)
		}
	}()

	return c
}

func ChannelDemos()  {
	var channels [10]chan<- int
	for i:=0;i<10;i++ {
		channels[i] = CreateWorker(i)
	}

	for i:=0;i<10;i++ {
		channels[i] <- 'a'+i
	}
}

func BufferedChannel()  {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
}
