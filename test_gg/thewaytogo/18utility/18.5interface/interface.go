package main

import (
	"fmt"
	"log"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

type Men interface {
	SayHi()
}

func main() {
	mike := Human{"Mike", 25, "222-222-XXX"}
	var v Men
	v = mike
	if v, ok := v.(Men); ok {
		v.SayHi()
	}
}

// 使用接口 实现一个类型分类函数
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}

// recover 终止panic

func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
