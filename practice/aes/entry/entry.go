package main

import (
	"fmt"
	"learnGo/practice/aes"
)

func main() {
	appSecret := []byte("A6DNS5iJechhRt4OO97Kz9cRlOWgoQsJ1564993030\n")
	iv := []byte("shulvlym12345678")
	originData := []byte(`{
	"detail_data":[
        {
            "owner_customer_code":"2",
            "invoice_type":"7",
            "pay_channel":"2",
            "taxpayer_type":"1032",
            "income_source_id":"20200927001",
            "income_account":"13213793819",
            "name":"陈兴生",
            "phone":"13213793819",
            "id_card":"420116199612081418",
            "pay_money":"8.88",
            "notify_url":"http://localhost:10610/api/home/back",
            "remark":"",
            "bank_digest":""
        },{
            "owner_customer_code":"2",
            "invoice_type":"7",
            "pay_channel":"2",
            "taxpayer_type":"1032",
            "income_source_id":"20200927002",
            "income_account":"13213793819",
            "name":"陈兴生",
            "phone":"13213793819",
            "id_card":"420116199612081418",
            "pay_money":"8.88",
            "notify_url":"http://localhost:10610/api/home/back",
            "remark":"",
            "bank_digest":""
        }
    ]
}`)

	result := aes.AesCBCEncrypt(appSecret, originData, iv)
	fmt.Println(result)

	//var aeskey = []byte("5McUOZkZMoJym0JN")
	//pass := []byte(`{"IP": "127.0.0.1", "name": "SKY"}`)
	//xpass, err := aes.AesEncrypt(pass, aeskey)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Printf("%v\n", xpass)
	////pass64 := base64.StdEncoding.EncodeToString(xpass)
	////fmt.Printf("加密后:%v\n",pass64)
	//
	//iv := aeskey
	//fmt.Println("IV：", iv)
	//var buffer bytes.Buffer
	//buffer.Write(iv)
	//buffer.Write(xpass)
	//realData := buffer.Bytes()
	//fmt.Println("拼接后的结果：", realData)
	//pass641 := base64.StdEncoding.EncodeToString(realData)
	//fmt.Println("想要的加密结果", pass641)

	//bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//tpass, err := aes.AesDecrypt(bytesPass, aeskey)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("解密后:%s\n", tpass)
}


