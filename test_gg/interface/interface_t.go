package main

import "fmt"

type A interface {
	Say()
}
type B interface {
	Hello()
}
type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("mmm hello")
}
func (m Monster) Say() {
	fmt.Println("mmm Say")
}

func Main() {
	var m Monster
	var a A = m
	var b B = m
	a.Say()
	b.Hello()
}
