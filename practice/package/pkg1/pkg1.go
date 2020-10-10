package pkg1

import (
	"fmt"
	"learnGo/practice/package/pkg2"
)

func ExecPkg1()  {
	fmt.Println("Exec Pkg1")
	pkg2.ExecPkg2()
}

func init()  {
	fmt.Println("pkg1 init")
}
