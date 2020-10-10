package main

import (
	"learnGo/practice/package/factory/base"
	_ "learnGo/practice/package/factory/cls1"
	_ "learnGo/practice/package/factory/cls2"
)

func main() {
	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}
