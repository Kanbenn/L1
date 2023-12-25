package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int)
	arr := []interface{}{1, "string", true, 0.5, ch}
	for _, i := range arr {
		fmt.Println(reflect.TypeOf(i))
		switch i.(type) {
		case int:
			fmt.Println("Int")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Bool")
		case float64:
			fmt.Println("Float64")
		case chan int:
			fmt.Println("Chan int")
		}
	}
}
