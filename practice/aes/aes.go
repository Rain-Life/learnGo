package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	fmt.Println("原始数据", string(origData[:]))
	block, err := aes.NewCipher(key)
	//fmt.Println("block内容：", block)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	fmt.Println("blockSize值：", blockSize)

	origData = PKCS5Padding(origData, blockSize)
	fmt.Println("填充后数据", string(origData[:]))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	fmt.Println("通过iv处理加密后：", blockMode)
	crypted := make([]byte, len(origData))
	fmt.Println("crypted：", crypted)
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

//封装完整的GO实现aes加密
/**
 * GO语言实现AES加密
 * @param appSecret 秘钥(51这边提供的app_secret)
 * @param originData 加密内容（保证是json格式的）
 * @param iv 偏移量（可以是随机生成的16位）
 */
func AesCBCEncrypt(appSecret, originData, iv []byte) string {
	if len(string(appSecret)) < 16 {
		panic("秘钥长度不合法")
	}
	
	appSecret = appSecret[:16]
	block, err := aes.NewCipher(appSecret)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	afterPaddingData := PKCS5Padding(originData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)//iv为16位
	crypted := make([]byte, len(afterPaddingData))
	blockMode.CryptBlocks(crypted, afterPaddingData)

	var buffer bytes.Buffer
	buffer.Write(iv)
	buffer.Write(crypted)
	readyEncryptData := buffer.Bytes()

	entryResData := base64.StdEncoding.EncodeToString(readyEncryptData)

	return entryResData
}


