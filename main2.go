package main

import "fmt"

type Color int

func (c Color) String() string {
	switch c {
	case ColorBlack:
		return "BLACK"
	case ColorBlue:
		return "BLUE"
	case ColorPink:
		return "PINK"
	case ColorYellow:
		return "YELLOW"
	default:
		panic("invalid selection")
	}

}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main2() {
	fmt.Println("the color is: ", ColorPink)
}
