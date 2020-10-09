package main

import "fmt"

type Data struct {
	complex []int
	instance InnerData
	ptr *InnerData
}

type InnerData struct {
	a int
}

const (
	SecondPerMinute = 60
	SecondPerHour = SecondPerMinute*60
	SecondPerDay = SecondPerHour*24
)

func funcTest1(a,b int) (int, int)  {
	return a/b,a%b
}

func resolveTime(second int) (minute, hour, day int) {
	day = second / SecondPerDay
	hour = second/ SecondPerHour
	minute = second/SecondPerMinute

	return
}

func passByValue(inFunc Data) Data {
	fmt.Printf("in Func value:%+v\n", inFunc)
	fmt.Printf("in Func ptr:%p\n", &inFunc)
	fmt.Println()
	return inFunc
}

func fire() {
	fmt.Println("fire")
}

func main() {
	r1,r2 := funcTest1(13, 3)
	fmt.Println(r1, r2)

	fmt.Println(resolveTime(1000))
	fmt.Println(resolveTime(18000))
	_,_,day := resolveTime(90000)
	fmt.Println(day)

	in := Data{
		complex: []int{1,2,3},
		instance: InnerData{5,},
		ptr: &InnerData{1},
	}

	fmt.Printf("in value:%+v\n", in)
	fmt.Printf("in ptr:%p", &in)
	fmt.Println()
	out := passByValue(in)
	fmt.Printf("in value:%+v\n", out)
	fmt.Printf("in ptr:%p", &out)

	fmt.Println()
	var f func()
	f = fire
	f()
}
