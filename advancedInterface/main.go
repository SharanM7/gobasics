package main

import (
	"fmt"
	"strings"
)

// type Storer interface {
// 	Get(int) (any, error)
// 	Put(any) (int, error)
// }

// type Store1 struct {
// }

// func (s1 Store1) Put(val any) (int, error) {
// 	log.Println("in put of store1")
// 	return 0, nil
// }

// func (s1 Store1) Get(id int) (any, error) {
// 	log.Println("in get of store1")
// 	return nil, nil
// }

type Request struct {
	version int
	s       string
	fn      FormatLogic
	Prefix  string
}

func main() {
	// fmt.Println("start")
	// s := Store1{}

	// s.Put(12)
	// s.Get(12)
	request := &Request{
		version: 3,
		s:       "asd",
		Prefix:  "P_",
	}

	request.Choice()

	fmt.Println("formatted :", Formatter(request.s, request.fn))
}

func (r *Request) Choice() {
	switch r.version {
	case 1:
		r.fn = FormatLogic1
	case 2:
		r.fn = FormatLogic2
	case 3:
		r.fn = FormatLogic3(r.Prefix)
	}
}

type FormatLogic func(string) string

func FormatLogic1(s string) string {
	return strings.ToUpper(s)
}

func FormatLogic2(s string) string {
	return strings.ToUpper(s) + ".csv"
}

func Formatter(s string, fn FormatLogic) string {
	return fn(s)
}

func FormatLogic3(prefix string) FormatLogic {
	return func(s string) string {
		return prefix + s
	}
}
