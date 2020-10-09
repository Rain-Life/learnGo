package main

import "fmt"

type Walkable struct {}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

type Flying struct {}

func (f *Flying) Fly()  {
	fmt.Println("can fly")
}

type Human struct {
	Walkable
}

type Bird struct{
	Walkable
	Flying
}

func main() {
	b := new(Bird)
	fmt.Println("Bird:")
	b.Walk()
	b.Fly()

	h := new(Human)
	fmt.Println("Human:")
	h.Walk()
}
