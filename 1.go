package main

import "fmt"

type Human struct {
	Action
}

type Action struct {
	counter int
}

func (a *Action) Eat() {
	a.counter -= 1
	fmt.Printf("Call Eat %d\n", a.counter)
}

func (a *Action) Shit() {
	a.counter++
	fmt.Printf("Call shit\n %d\n", a.counter)
}
