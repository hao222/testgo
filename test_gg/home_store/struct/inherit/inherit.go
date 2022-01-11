package main

import "fmt"

// go中的继承实现  一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法， 从而实现了继承特性

type Stu struct {
	Name  string
	Age   int
	Score int
}

func (stu *Stu) ShowInfo() {
	fmt.Printf("%v--%v---%v", stu.Name, stu.Age, stu.Score)
}

// 继承自Stu 可以使用Stu的showinfo 方法
type Puplie struct {
	Stu
}

func (p *Puplie) tesing() {
	fmt.Println("这是Puplie的私有方法")
}

func main() {
	var pp = Puplie{}
	pp.Stu.Name = "xxx"
	pp.Stu.Age = 12
	pp.Score = 121
	pp.tesing()
	pp.Stu.ShowInfo()
}
