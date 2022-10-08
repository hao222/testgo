package main

import (
	"fmt"
	"test_gg/test_gg/thewaytogo/10struct_method/6method/person"
)

type TwoInts struct {
	a int
	b int
}

// 指针方法和值方法都可以在指针或非指针上被调用
type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

// 在类型中嵌入功能
// A： 聚合/组合：包含一个所需功能类型的具名字段
// B： 内嵌：内嵌（匿名地）所需功能类型

// 方式A：
type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

// 方式B：
//func (l *Log) Add(s string) {
//	l.msg += "\n" + s
//}
//
//func (l *Log) String() string {
//	return l.msg
//}
//
//func (c *Customer) String() string {
//	return c.Name + "\nLog:" + fmt.Sprintln(c.Log.String())
//}

type Base struct{}

func (Base) Magic() {
	fmt.Println("14.1base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))
	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)

	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
	println()
	p := new(person.Person)
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())

	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
// 接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，
// 可以是 int、bool、string 或数组的别名类型。但是接收者不能是一个接口类型
func (th *TwoInts) AddThem() int {
	return th.a + th.b
}

func (th *TwoInts) AddToParam(param int) int {
	return th.a + th.b + param
}

//类型和作用在它上面定义的方法必须在同一个包里定义

// 多重继承：
type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

// CameraPhone 同时 继承Camera Phone 并使用其中的方法
type CameraPhone struct {
	Camera
	Phone
}
