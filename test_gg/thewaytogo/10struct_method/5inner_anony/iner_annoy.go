package main

import "fmt"

// 匿名字段 以及  内嵌结构体

type innerS struct {
	one int
	two int
}

type outerS struct {
	b int
	c float32
	int
	innerS // 匿名字段结构体    另类的继承
}

// 避免内嵌结构体字段命名冲突
type A struct{ a int }
type B struct{ a, b int }

type C struct {
	A
	B
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.one = 5
	outer.two = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.one)
	fmt.Printf("outer.in2 is: %d\n", outer.two)

	outer2 := outerS{6, 7.7, 80, innerS{2, 3}}
	fmt.Println(outer2)
}
