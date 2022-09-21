package main

import (
	"fmt"
	"math"
)

// 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量
//type Namer interface {
//	Method1(param_list) return_type
//	Method2(param_list) return_type
//	...
//}

// 接口可以嵌套接口
type ReadWrite interface {
	Read() bool
	Write() bool
}
type File interface {
	ReadWrite
	Close()
}

// 如果 Shaper 有另外一个方法 Perimeter()，但是 Square 没有实现它，即使没有人在 Square 实例上调用这个方法，编译器也会给出上面同样的错误

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := &Circle{2.5}
	fmt.Println("Looping through shapes for area ...")
	// shapes := []Shaper{Shaper(r), Shaper(q), Shaper(c)}
	shapes := []Shaper{r, q, c}
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
