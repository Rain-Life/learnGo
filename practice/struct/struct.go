package main

import (
	"fmt"
	"math"
)

type Bag struct {
	items []int
}

type Property struct {
	value int
}

type Vec2 struct {
	X,Y float32
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2  {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx * dx + dy * dy)))
}

func (v Vec2) Normalize() Vec2  {
	mag := v.X * v.X + v.Y * v.Y
	if mag > 0 {
		oneOverMag := 1/float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}

	return Vec2{0, 0}
}


func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

func (b *Bag) Insert1(itemId int)  {
	b.items = append(b.items, itemId)
}

func (p *Property) setValue(v int)  {
	p.value = v
}

func (p *Property) Value() int  {
	return p.value
}

type myInt int

func (m myInt) isZero() bool {
	return m==0
}

func (m myInt) Add(other int) int {
	return int(m) + other
}

func main() {
	//var bag Bag
	////bag := new(Bag)
	//Insert(&bag, 1001)
	//bag.Insert1(1002)
	//
	//p := new(Property)
	//p.setValue(10)
	//println(p.Value())

	var b myInt
	fmt.Println(b.isZero())

	b = 1
	fmt.Println(b.Add(2))
}
