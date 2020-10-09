package main

import dictionary "learnGo/practice/dictionary/switch"

func main() {
	//map.go
	//dict := dictionary.NewDictionary()
	//
	//dict.Set("MyFactory", 60)
	//dict.Set("TerraCraft", 36)
	//dict.Set("Don'tHungry", 24)
	//
	//favorite := dict.Get("TerraCraft")
	//fmt.Println("favorite", favorite)
	//
	//dict.Visit(func(k, v interface{}) bool {
	//	if v.(int) > 40 {
	//		fmt.Println(k, "Is expensive")
	//		return true
	//	}
	//
	//	fmt.Println(k, "Is cheap")
	//	return true
	//})

	//switch.go
	dictionary.PrintType(1024)
	dictionary.PrintType("pid")
	dictionary.PrintType(true)

	//switch2.go
	dictionary.Print(new(dictionary.Alipay))
	dictionary.Print(new(dictionary.Crash))
}
