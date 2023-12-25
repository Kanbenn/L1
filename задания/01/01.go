package main

import (
	"fmt"
)

func main() {
	var a Action
	a.Walk()
}

type Human struct{}

func (h *Human) Walk() {
	fmt.Println("Human Walk")
}

type Action struct {
	h Human
}

func (a *Action) Walk() {
	fmt.Println("Action Walk")
	a.h.Walk()
}
