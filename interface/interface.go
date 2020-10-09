package learn_interface

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type File struct {

}

func (d *File) WriteData(data interface{}) error {
	fmt.Println("Write Data:", data)
	return nil
}