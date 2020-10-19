package main

import (
	"fmt"
	"regexp"
)

const text = "My email is shulv@gmail.com"
const text1 = `My email is shulv@gmail.com
email1 is abc@def.org
email2 is    kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)//.意思是匹配任何一个字符，+是指匹配一个或多个，*是匹配0个或多个
	match := re.FindString(text)
	fmt.Println(match)
	regex1()
	regex2()
}

func regex1() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	match := re.FindAllString(text1, -1)//-1指的是匹配所有的
	fmt.Println(match)
}

//提取出邮件的名字以及后缀这些
func regex2() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text1, -1) //会按照正则中的括号中进行提取
	for _,email := range match {
		fmt.Println(email)
	}
}
