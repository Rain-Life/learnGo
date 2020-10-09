package main

import (
	"encoding/base64"
	"fmt"
	"learnGo/practice/aes"
)

func main() {
	var aeskey = []byte("321423u9y8d2fwfl")
	pass := []byte("vdncloud123456")
	xpass, err := aes.AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n",pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := aes.AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}

