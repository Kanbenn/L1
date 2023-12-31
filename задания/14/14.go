package main

import (
	"fmt"
	"reflect"
)

type storer interface {
	Lol()
}

type loler struct{}

func (l loler) Lol() { fmt.Println("Lol") }

func main() {
	ch := make(chan int)
	lol := loler{}
	var st = storer(lol)
	arr := []any{1, "string", true, 0.5, ch, lol, st}
	for _, i := range arr {
		switch i.(type) {
		case int:
			fmt.Println("Int")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Bool")
		case float64:
			fmt.Println("Float64")
		default:
			fmt.Println("unexpected type:", reflect.TypeOf(i))
		}
	}
}
