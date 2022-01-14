package main

import (
	"fmt"
	"reflect"
)

// 常量只能 数值 string bool
const nums = 111

// 反射   对结构体序列化 json `name`  // 编写函数的适配器， 桥连接
// reflect包实现了运行时反射，允许程序操作任意类型的对象。典型用法是用静态类型interface{}保存一个值，通过调用TypeOf获取其动态类型信息，
// 该函数返回一个Type类型值。调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。
//反射在运行时动态获取变量的各种信息， 比如变量的类型type， 类别kind。 如果是结构体变量， 可以获取到结构体本身的信息（包括结构体的字段， 方法）
// 通过反射， 可以修改变量的值，可以调用关联的方法
// reflect.TypeOf(变量名)， 获取变量的类型， 返回reflect.Type类型
// reflect.ValueOf(), 获取变量的值， 返回reflect.Value类型 是一个结构体类型
// 变量， interface{} reflect.Value 是可以相互转换的
//func test(b interface{})  {
//	var stu Stu
//	// 将interface{} 转成 reflect.Value
//	rVal:=reflect.ValueOf(b)
//	// 将 reflect.Value --> interface{}
//	iVal:=rVal.Interface()
//	// 将interface{} 转成原来的变量类型， 使用类型断言
//	v:=iVal.(struct)
//}

// int reflect
func reflectcase01(b interface{}) {
	// 反射获取传入变量的type kind 值
	rType := reflect.TypeOf(b)
	fmt.Println("rtype=", rType, rType.Name())
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v rval type=%T\n", rValue, rValue)

	//
	n2 := 2 + rValue.Int()
	fmt.Println("n2---", n2)
	iv := rValue.Interface()
	num2 := iv.(int)
	fmt.Println("num2= ", num2)
}

// struct reflect
func reflectcase02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rtype=", rType, rType.Name())
	rValue := reflect.ValueOf(b)
	// kind 代表Type类型值表示的具体分类， 零值表示非法分类
	fmt.Printf("rValue=%v rval type=%T rval kind=%v\n", rValue, rValue, rValue.Kind())
	iv := rValue.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iv, iv)
	// 将interface 通过断言转成需要的类型
	// 使用switc 控制 类型返回数据
	stu, ok := iv.(Stu)
	if ok {
		fmt.Printf("stu.Name= %v \n", stu.Name)
	}
}

type Stu struct {
	Name string
	Age  int
}

func reflectupdate(b interface{}) {
	rVal := reflect.ValueOf(b)
	// 此时的kind  为指针类型 需要 转换 获取到值 使用Elem
	rVal.Elem().SetInt(20)
}

func main() {
	var num = 100
	reflectcase01(num)
	stu := Stu{
		Name: "aaa",
		Age:  20,
	}
	reflectcase02(stu)
	var num1 = 100
	reflectupdate(&num1)
	fmt.Println("num1 ----", num1)
}

// reflect.Value.Kind  获取变量的类别， 返回一个常量 类型  reflect.Type.Kind
// Type是类型 Kind是类别(具体类型)， 两者有可能相同也有可能不同  stu Stu Type：函数.Stu  Kind : struct
// 通过reflect.Value 取值的基本数据类型 需要和原来数据类型一致 reflect.Value(x).Int()  x为int类型
// 通过反射修改变量 SetXxx需要通过对应的指针类型来完成， 这样才能改变传入的变量的值， 也会用到reflect.Value.Elem()
