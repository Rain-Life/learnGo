package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func padding(src []byte,blocksize int) []byte {
	padnum:=blocksize-len(src)%blocksize
	pad:=bytes.Repeat([]byte{byte(padnum)},padnum)
	return append(src,pad...)
}

func unpadding(src []byte) []byte {
	n:=len(src)
	unpadnum:=int(src[n-1])
	return src[:n-unpadnum]
}

func encryptAES(src []byte,key []byte) []byte {
	block,_:=aes.NewCipher(key)
	src=padding(src,block.BlockSize())
	blockmode:=cipher.NewCBCEncrypter(block,key)
	blockmode.CryptBlocks(src,src)
	return src
}

func decryptAES(src []byte,key []byte) []byte {
	block,_:=aes.NewCipher(key)
	blockmode:=cipher.NewCBCDecrypter(block,key)
	blockmode.CryptBlocks(src,src)
	src=unpadding(src)
	return src
}

func main()  {
	x:=[]byte("世界上最邪恶最专制的现代奴隶制国家--朝鲜")
	key:=[]byte("hgfedcba87654321")
	x1:=encryptAES(x,key)
	x2:=decryptAES(x1,key)
	fmt.Print(string(x2))
}