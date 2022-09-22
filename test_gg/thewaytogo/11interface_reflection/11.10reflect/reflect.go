package main

import (
	"fmt"
	"reflect"
)

// 反射是用程序检查其所拥有的结构，尤其是类型的一种能力
// 这是元编程的一种形式。反射可以在运行时检查类型和变量，例如：它的大小、它的方法以及它能“动态地”调用这些方法
// reflect.TypeOf 和 reflect.ValueOf，返回被检查对象的类型和值

func RE() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

	// 是否可设置是 Value 的一个属性，并且不是所有的反射值都有这个属性：可以使用 CanSet() 方法测试是否可设置
	// 要想 v 的更改能作用到 x，那就必须传递 x 的地址
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	// v 现在的类型是 *float64 并且仍然是不可设置的
	fmt.Println("settability of v:", v.CanSet())
	// elem 函数为可设置
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	//v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
}
