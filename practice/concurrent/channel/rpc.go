package channel

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string, error) {
	//向服务器发送请求
	ch <- req

	//等待服务器返回
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "",errors.New("Time Out")
	}
}

func RPCServer(ch chan string)  {
	for  {
		data := <-ch
		fmt.Println("Server Received:", data)
		//模拟超时
		//time.Sleep(time.Second*2)
		ch <- "roger"//告诉客户端，我收到数据了
	}
}

func Transfer()  {
	ch := make(chan string)

	go RPCServer(ch)

	recv, err := RPCClient(ch, "Hello Shulv")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received:", recv)
	}
}
