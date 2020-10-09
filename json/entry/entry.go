package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

var str = []byte(`
{"code":10011,"message":"\u8fd8\u662f\u6d4b\u8bd5","data":{"key1":[{"key1_1":"value1_1","key1_2":"value1_2","key1_3":"value1_3"},{"key1_1":"value1_1","key1_2":"value1_2","key1_3":"value1_3"}],"key2":{"key2_1":"value2_1","key2_2":"value2_2","key2_3":"value2_3"}}}`)

func main() {
	resp := Response{}
	err := json.Unmarshal(str, &resp)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp.Data["key1"])
	key1 := resp.Data["key1"].([]interface{})
	fmt.Printf("%T",key1)
	for k,v := range key1 {
		fmt.Println(k, v.(map[string]interface{})["key1_1"])
	}
}
