package main

import "fmt"

/*
数组是可变的 相同唯一类型

把一个大数组传递给函数会消耗很多内存
有两种方法可以避免这种情况：
	传递数组的指针
	使用数组的切片
*/

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

//var screen [WIDTH][HEIGHT]pixel // 多维数组

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func main() {

	//a := [...]string{"a", "b", "c", "d"}
	//for i := range a {
	//	fmt.Println("Array item", i, "is", a[i])
	//}

	var arr1 = new([5]int) // 用new 创建类型 *[5]int
	//var arr2 [5]int    // 类型是 [5]int
	for i := range arr1 {
		fmt.Println("Array item", i, "is", arr1[i])
	}
	// 改变原数组
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar

	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
