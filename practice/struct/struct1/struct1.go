package main

import "fmt"

type BasicColor struct {
	R,G,B float32
}

type Color struct {
	Basic BasicColor
	Alpha float32
}

func main() {
	var c Color
	c.Basic.R = 1
	c.Basic.G = 1
	c.Basic.B = 0

	c.Alpha = 1
	fmt.Printf("%+v", c)
}
