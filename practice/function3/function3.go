package main

import (
	"bytes"
	"fmt"
)

func joinString(slist ...string) string {
	var b bytes.Buffer
	for _,s := range slist {
		b.WriteString(s)
	}

	return b.String()
}

func printTypeValue(slist ...interface{}) string  {
	var b bytes.Buffer
	for _,s := range slist {
		str := fmt.Sprintf("%v", s)
		var typeString string
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		b.WriteString("Value：")
		b.WriteString(str)
		b.WriteString(" type：")
		b.WriteString(typeString)
		b.WriteString("\n")
	}

	return b.String()
}

func rawPrint(rawList ...interface{}) {
	for _,s := range rawList {
		fmt.Println(s)
	}
}

func print(slist ...interface{})  {
	rawPrint(slist...)
}

func main() {
	fmt.Println(joinString("shulv", " and ", "liangyu ", "love"))
	fmt.Println(printTypeValue(100, "shulv", true))
	print(1,2,3)
}
