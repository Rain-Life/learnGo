package main

import "fmt"

type Weapon int
type ChipType int

func constTest()  {
	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	fmt.Println(Arrow, Shuriken,SniperRifle, Rifle, Blower)
}

func consttest1()  {
	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)

	fmt.Printf("%d %d %d", FlagRed, FlagGreen, FlagBlue)
	fmt.Println()
	fmt.Printf("%b %b %b", FlagRed, FlagGreen, FlagBlue)
}
const (
	None ChipType = iota
	CPU
	GPU
)
func (c ChipType) String() string {
	switch c {
	case None:
		return "NOne"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {
	constTest()
	consttest1()
	fmt.Println()
	fmt.Printf("%s %d", CPU, CPU)
}
