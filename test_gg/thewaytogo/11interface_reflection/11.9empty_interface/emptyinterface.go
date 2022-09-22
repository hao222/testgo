package main

import "fmt"

// 空接口 类似一个基类 可以给一个空接口类型的变量 var val interface {} 赋任何类型的值

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", *t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}

	TypeSwitch()
}

// 空接口在 type-switch 中联合 lambda 函数用法

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)

	root := NewNode(nil, nil)
	root.SetData("root node")
	// make child (leaf) nodes:
	a := NewNode(nil, nil)
	a.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.le = a
	root.ri = b
	fmt.Printf("\n %v \n", root)
}

// 空接口 构建通用类型或包含不同类型变量的数组

type Element interface{}

type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}
func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

// 复制数据切片至空接口切片 必须使用for-range 语句一个一个的赋值， 直接相等赋值会报错 内存中的布局不一样
//var dataSlice []myType = FuncReturnSlice()
//var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
//for i, d := range dataSlice {
//interfaceSlice[i] = d
//}

// 通用类型的节点数据结构
// 二叉树的部分代码：通用定义、用于创建空节点的 NewNode 方法，及设置数据的 SetData 方法
type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

// 接口到接口  一个接口的值可以赋值给另一个接口变量，只要底层类型实现了必要的方法, 否则会报类型的错误
