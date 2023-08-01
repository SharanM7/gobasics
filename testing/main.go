package main

import (
	"log"
	"reflect"
)

func SumInt(xi ...int) int {
	c := 0
	for _, i := range xi {
		c += i
	}
	return c
}

type Req struct {
	I   int
	ids []int
}

func main() {
	r1 := Req{
		I:   1,
		ids: []int{1, 2, 3, 4, 5, 6, 7, 8},
	}
	r2 := Req{
		I:   1,
		ids: []int{1, 2, 3, 4, 5, 6, 7, 8},
	}

	if reflect.DeepEqual(r1, r2) {
		log.Println("matched")
	} else {
		log.Println("not matched")
	}

}
