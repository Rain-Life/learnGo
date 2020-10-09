package main

import (
	"bytes"
	"fmt"
	"strings"
)

func connectString() {
	hammer := "吃我一锤"
	sickle := "死吧"
	var stringBuilder bytes.Buffer

	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	fmt.Println(stringBuilder.String())
}

func main() {
	tracer := "死神来了，死神bye bye"
	comma := strings.Index(tracer, "，")
	fmt.Println(comma)
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])

	connectString()
}
