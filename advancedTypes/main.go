package main

import (
	"fmt"
)

type User struct {
	firstname string
	lastname  string
	age       int
	SpecialPower
}

type SpecialPower struct {
	speed int
}

func (u User) foo() {
	fmt.Println("here in foo of user")
}

// func (s SpecialPower) foo() {
// 	fmt.Println("here in foo of sp")
// }

type Color int

const (
	ColorBlack Color = iota + 1
	ColorBlue
	ColorRed
	ColorYellow
)

func (c Color) String() string {
	switch c {
	case ColorBlack:
		return "BLACK"
	case ColorBlue:
		return "BLUE"
	case ColorRed:
		return "RED"
	case ColorYellow:
		return "YELLOW"
	default:
		return "INVALID"
	}
}

func main() {
	// num := getNumber()
	// log.Println("here", num)
	u := User{
		firstname: "a",
		lastname:  "b",
		age:       30,
		SpecialPower: SpecialPower{
			speed: 100,
		},
	}
	u.speed = 500
	u.foo()
	fmt.Printf("data : %+v\n", u)

	fmt.Println("color :", ColorBlue)
	var k Color = 100
	fmt.Println("color var to string :", k)
}
