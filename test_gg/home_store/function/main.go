package main

import "fmt"

// 方法是作用在指定的数据类型上的， 因此自定义类型 都可以有方法
// struct里的方法又叫做行为

type A struct {
	Name string
}

// 调用机制和函数一样， 不同的是 变量调用方法时， 该变量本身也会作为一个参数传递到方法 如果是值类型，则进行值拷贝， 如果变量是引用类型， 则进行地址拷贝
func (a A) test() {
	fmt.Println("A 结构体的方法-", a.Name)
}

//func test(n1 int, name string) (int, bool) {
//	return 1, True
//}
// 结构体类型是值类型， 子方法调用中，遵守值类型传递机制， 是值拷贝传递方式
// 修改结构体变量的值， 可以通过结构体指针的方式处理
// 方法首字母大小写 访问规则

// 一个变量定义了string() 方法 fmt.println 则默认会调用string返回
func (a *A) String() string {
	return fmt.Sprintf("Name=[%v]", a.Name)
}

// 一个方法实现 +-*/
type Cal struct {
	Num1 float64
	Num2 float64
}

func (cal *Cal) getRes(ope byte) float64 {

	res := 0.0
	switch ope {
	case '+':
		res = cal.Num1 + cal.Num2
	default:
		fmt.Println("运算符有误")
	}
	return res

}

// 方法函数区别  1 调用方式不同  2 函数接收为值类型， 不能将指针类型数据直接传递， 接受者为指针类型也不能将值类型直接传递
// 方法， 接受者为指针或者值类型都可直接传递
func main() {
	var a A
	a = A{"xcc"}
	a.test()
	var aa A
	aa = A{"toms"}
	fmt.Println("-------", &aa)
}
