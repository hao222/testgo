package main

import (
	"fmt"
	"reflect"
)

// 定制类型： 通过类型别名和结构体形式
// 组成结构体类型的那些数据称为 字段 (fields)。每个字段都有一个类型和一个名字；在一个结构体中，字段名字必须是唯一的
// 字段可以是任意类型
// 结构体和它所包含的数据在内存中是以连续块的形式存在的

//type identifier struct {
//	filed1 type1
//	filed2 type2
//}

// 类型就是类
// 代码复用通过组合和委托实现 多态通过接口的使用来实现

type T struct{ a, b int }
type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	var s T
	s.a = 6
	s.b = 8
	fmt.Printf("%v", s)
	t := new(T) // t 通常被称做类型 T 的一个实例 (instance) 或对象 (object)
	fmt.Printf("%v", t)
	// 初始化一个结构体实例
	ms := struct1{10, 15.5, "Chris"}
	mss := struct1{i1: 5, f1: 1.1}
	fmt.Printf("%v-----%v", ms, mss)

}

// 链表
type Node1 struct {
	data float64
	su   *Node1
}

// 双向链表
type Node2 struct {
	pr   *Node2
	data float64
	su   *Node2
}

// 二叉树
type Tree struct {
	le   *Tree
	data float64
	ri   *Tree
}

// 使用工厂构建结构体实例
type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}

	//// 强制使用工厂方法
	//nf := new(File)
	//return nf
}

// tags 打标签 注释 通过reflect typeof 获取
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
