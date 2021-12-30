package main

import "fmt"

// struct， go没有class 但是可以围绕struct进行类的操作 有封装， 继承和多肽通过 你名字短或者接口来实现
// OOP面向对象  面向接口编程

// 指针 slice map的零值都是nil  即没有分配内存空间
type Person struct {
	Name  string            // ""
	Age   int               // 0
	Score [5]float64        // [0,0,0,0,0]
	Ptr   *int              // nil
	slice []int             // [] nil
	map1  map[string]string // map[]  nil
}

type S struct {
	Name string
}

func main() {
	// 1
	var p1 Person
	// 切片 map ptr 需要make分配空间
	p1.slice = make([]int, 5)
	p1.slice[0] = 100
	fmt.Println("struct -----", p1)

	var ss S
	ss.Name = "xxx"
	sss := ss //  将结构体实例ss 赋给sss， 修改sss 并不会影响ss  值类型
	sss.Name = "zzz"
	fmt.Println(ss, sss)

	// 2
	p2 := Person{}
	p2.Name = "rrr"
	p2.Age = 120
	fmt.Println(p2)

	// 3 指针形式
	//p3 := new(Person)
	var p3 *Person = new(Person)
	(*p3).Name = "smith"
	p3.Age = 12 // 可以直接使用.操作， 内部处理
	fmt.Println(*p3)
	// 4
	var p4 *Person = &Person{}
	(*p4).Name = "scott"
	p4.Age = 34
	fmt.Println(*p4)
}

// 1 结构体的变量在内存中是连续的
// 2 结构体为用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字 个数 类型）

type Point struct {
	x string
	y string
}

// 地址在内存中是连续分布的
type Rect struct {
	left, right Point
}

// 有两个*point指针类型， 这两个*point类型的本身地址是连续的， 但是他们指向的地址不一定连续
type Rect1 struct {
	left, right *Point
}
