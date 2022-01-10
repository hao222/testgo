package main

import "fmt"

// 结构体可以使用嵌套匿名结构体的所有字段和方法  即 首字母大写或小写的字段和方法都可以使用
type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A say ok", a.Name)
}
func (a *A) hello() {
	fmt.Println("A say ok", a.age)
}

func (b *B) hello() {
	fmt.Println("B hello", b.age)
}

// 当结构体和匿名结构体有相同的字段或者方法时， 编译器会采用就近访问原则  如希望访问匿名结构体的方法或者字段， 可以通过匿名结构体名来区分

type B struct {
	A
	age int
}

// 结构体嵌入两个匿名结构体， 如两个匿名结构体有相同的字段和方法（结构图本身没有同名的字段和方法， 就必须明确指定匿名结构体名， 否则会报错
type C struct {
	A
	B
}

// 如果一个结构体嵌套一个有名的结构体， 这种模式称为组合， 在访问组合的结构体的字段或方法时， 必须带上结构体的名字
type D struct {
	a A // 有名结构体
}

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var b B
	b.A.Name = "tom"
	// 匿名结构体字段访问可以简化
	b.age = 10
	b.A.age = 19
	b.A.SayOk()
	b.hello()
	b.A.hello()

	var d D
	//d.Name = "xxx"   // 报错 d中 a为有名结构体
	d.a.Name = "xxx"

	// 嵌套匿名结构体后， 也可以再创建结构体变量时， 直接指定各个匿名结构体字段的值
	tv1 := TV{
		Goods{
			Name:  "xxx",
			Price: 12.12,
		},
		Brand{
			Name: "ner-1",
		},
	}
	fmt.Println(tv1)

	tv := TV{Goods{"tvoo1", 5000.11}, Brand{"haier--1", "shandong"}}
	fmt.Println(tv)
	// 嵌套指针类型的匿名结构体
	tv2 := TV2{&Goods{"tvoo1", 5000.11}, &Brand{"haier--1", "shandong"}}
	fmt.Println(*tv2.Goods, *tv2.Brand)

}
