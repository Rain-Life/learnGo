package main

import "net/http"

func main() {
	resp, err :=http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}


}
