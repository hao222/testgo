package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	// 由于接口是一般类型， 不知道具体类型， 如果要转成具体类型， 需要使用类型断言
	var a interface{}
	var point Point = Point{1, 2}
	a = point // 空接口可以接收任意类型 确保原来的空接口指向的就是断言的类型， 否则会报错panic
	//var b Point
	//	类型断言 + 断言检查
	b, ok := a.(Point)
	if ok {
		fmt.Println("------", b)
	} else {
		fmt.Println("convert fail")
	}

}
