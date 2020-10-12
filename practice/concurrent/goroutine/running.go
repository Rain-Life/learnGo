package concurrent

import (
	"fmt"
	"time"
)

func Running()  {
	var times int
	for {
		times++
		fmt.Println("tick ", times)

		time.Sleep(time.Second)
	}
}

func Run()  {
	go Running()

	var input string
	fmt.Scanln(&input)
}
