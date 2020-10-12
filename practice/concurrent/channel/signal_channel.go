package channel

func Test() {
	ch := make(chan int)
	var chSendOnly chan<- int = ch
	var chRecvOnly <-chan int = ch

	ch1 := make(<-chan int)
	che2 := make(chan<- int)
}
