package main

import (
	"fmt"
	"strings"
	"time"
)

// 值传递  引用传递
//指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。所以，按引用传递也是按值传递
//函数调用时，像切片 (slice)、字典 (map)、接口 (interface)、通道 (channel) 这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）
//尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂
// 使用指针传参 更改原有变量值

var num = 10
var numx2, numx3 int

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	// var i1 int = MultiPly3Nums(2, 5, 6)
	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
	// 空白符用来匹配一些不需要的值，然后丢弃掉
	numx2, numx3, _ = getX2AndX3_2(num)
	println("-----", numx2, numx3)
	printInts(2, 3)
	println()
	printInts(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	println()
	slice := []interface{}{11, 2, 3, 44}
	show(slice...)
	function1()
	println()

	a()
	println()

	callback(11, Add)

	asciiOnly := func(c rune) rune {
		if c > 127 {
			return '1'
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))

	// 匿名函数也为闭包，关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值
	// 匿名函数还可以配合 go 关键字来作为 goroutine 使用
	// 闭包经常被用作包装函数：它们会预先定义好 1 个或多个参数以用于包装。另一个不错的应用就是使用闭包来完成更加简洁的错误检查
	xx := func(x, y int) int { return x + y }(3, 4)
	println("匿名函数：----", xx)

	start := time.Now()
	Add(1, 3)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Add took this amount of time: %s\n", delta)

}

func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

// 非命名返回值  命名返回值
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 命名返回值
func getX2AndX3_2(input int) (x2 int, x3 int, x4 string) {
	x2 = 2 * input
	x3 = 3 * input
	x4 = ""
	// return x2, x3
	return
}

// 变长参数 采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变参函数
func printInts(list ...int) {
	//	for i:=0; i<len(list); i++ {
	//		fmt.Printf("%d\n", list[i])
	//	}
	for _, v := range list {
		fmt.Printf("%d\n", v)
	}
}

// 如果变长参数的类型并不是都相同的呢   使用结构体或者使用空接口接收
func show(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

type Options struct {
	par1 int
	par2 string
}

// ------------------------------defer
// defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
// 关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。
// 关闭文件流 解锁一个加锁的资源 打印最终报告 关闭数据库连接 实现代码追踪 记录函数的参数与返回值
func function1() {
	defer function2()

	fmt.Printf("In function1 at the top\n")
	fmt.Printf("In function1 at the bottom!\n")
}
func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling 6function!")
}

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func a() {
	i := 0
	//defer fmt.Println(i)   // 打印0
	i++
	defer fmt.Println(i) // 打印 1
	return
}

// ------------------------------ 内置函数 ----
/*
close() 管道通信
len() 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；
cap() 是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道，不能用于 map）
new() 和 make() 均是用于分配内存：
new() 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）
copy()、append() 用于复制和连接切片
panic()、recover() 两者均用于错误处理机制
print()、println() 在部署环境中建议使用 fmt 包
complex()、real ()、imag()  用于创建和操作复数
*/

// --------------------- 递归函数 当一个函数在其函数体内调用自身，则称之为递归 ---------------
// 斐波那契数列 返回值和返回下标
func fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci(n - 1)
		v2, _ := fibonacci(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}

// 可以使用相互调用的递归函数：多个函数之间相互调用形成闭环
// 前30个阶乘
func Factorial(n uint64) (fac uint64) {
	fac = 1
	if n > 0 {
		fac = n * Factorial(n-1)
		return
	}
	return
}

// -------------------- 函数作为参数
// 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

// -----------------
