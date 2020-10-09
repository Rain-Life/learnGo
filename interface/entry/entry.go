package main

import (
	"io"
	learn_interface "learnGo/interface"
)

func main() {
	f := new(learn_interface.File)
	var writer learn_interface.DataWriter

	writer = f
	writer.WriteData("data")

	//interface3
	var wc io.WriteCloser = new(learn_interface.Device)
	wc.Write(nil)
	wc.Close()
}
