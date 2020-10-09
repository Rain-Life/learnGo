package main

import "fmt"

var eventByName = make(map[string][]func(interface{}))

type Actor struct {
}

func RegisterEvent(name string, callback func(interface{}))  {
	list := eventByName[name]
	list = append(list, callback)
	eventByName[name] = list
}

func CallEvent(name string, param interface{}) {
	list := eventByName[name]
	for _, callback := range list {
		callback(param)
	}
}

func (a *Actor) OnEvent(param interface{})  {
	fmt.Println("actor event", param)
}

func GlobalEvent(param interface{})  {
	fmt.Println("global event", param)
}

func main() {
	a := new(Actor)
	RegisterEvent("OnSkill", a.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	CallEvent("OnSkill", 100)
}
