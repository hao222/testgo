package main

// go 定义的变量一定要使用 不使用会报错

import (
	"fmt"
	"math"
	"math/cmplx"
	"sync"
)

var wg sync.WaitGroup

var aa = 3
var (
	cc = 22
	bb = 233
) // 统一定义
// ss := true  不能使用 需要var 定义变量

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q", a, s)
}

func test() {
	var a int = 3
	var s string = "abc" // 初始值
	var c = 333
	var d, e, f = "dddd", true, 32 // 1  可以自动决定类型
	dd, ee, ff := "dddd", true, 32 // 2 第一次变量需要：冒号  后面再修改的时候不需要
	dd = "haha"
	fmt.Printf("\n %d %q", a, s)
	fmt.Println(c, d, e, f)
	fmt.Println(dd, ee, ff)
}

func euler() {
	c := 3 + 4i

	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%0.3f", cmplx.Exp(1i*math.Pi)+1) // 欧拉公式
}

func triangle() {
	var a, b int = 3, 4
	var c int
	var d int
	var e float64 = 2.99999
	c = int(math.Sqrt(float64(a*a + b*b))) // 类型转换是 强制类型
	d = int(math.Ceil(e))
	fmt.Println(c, d)
}

func consts() {
	const filename = "aaa.txt"
	const a, b = 3, 5
	var c int
	c = int(math.Sqrt(a*a + b*b)) // 不定义 a，b 则是文本 不需要转类型
	fmt.Println(c, filename)
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		py     = 2
		golang = 3
	)

	const (
		a = iota // 实现自增 枚举
		b
		c
	)
	const (
		bb = 1 << (10 * iota)
		kb
		mb
	)
	fmt.Println(a, b, c)
	fmt.Println(cpp, golang, py)
	fmt.Println(bb, kb, mb)
}

func sums(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum

}

func AddUpper() func(int) int {
	n := 10
	return func(i int) int {
		n += i
		return n
	}

}

func main() {
	// go run 编译执行放在一起  类似一个执行脚本文件的形式
	// 一般是先编译 go build  生成.exe文件之后 再执行.exe
	// 先编译.exe文件可以放在没有go环境的机器上运行，但是文件会很大， 编译器会将程序运行依赖的库文件包含在可执行文件中
	/*
		go 的执行入口是main() 函数
		go 严格区分大小写
		go 方法由一条条语句构成， 每个语句后面不需要分号（自动在每行后加分号）
		go 编译是一行行编译， 因此我们一行就写一条语句， 不能把多条语句写在同一个， 否则报错
		go 定义的变量或者import的包如果没有使用到， 代码不能编译成功
		大括号成对出现的， 缺一不可

	*/
	variable()
	test()
	fmt.Println(aa, bb, cc)
	euler()
	triangle()
	consts()
	enums()
	fmt.Println("天龙潭日好风光无限天\r张三丰") // \r 回车 后面的会覆盖前面的值

	nu := 10
	switch nu {
	case 10:
		print("ok1")
		fallthrough // 会自动执行下一个case 无需判断
	case 20:
		print("ok2")

	case 30:
		print("ok3")
	default:
		print("未找到")
	}

	var xx interface{}
	var y = 10.0
	xx = y
	switch i := xx.(type) {
	case nil:
		fmt.Printf("x 类型 %T", i)
	case int:
		fmt.Printf("x shi int")
	case float64:
		fmt.Printf("x shi float")

	}
	sr := "hello world北京" // 出现乱码
	for i := 1; i < len(sr); i++ {
		fmt.Printf("%c\n", sr[i])
	}

	for index, val := range sr {
		fmt.Printf("index= %d, val=%c", index, val)
	}

	str2 := []rune(sr)

	for i := 1; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	res := sums(1, 2, 3, 4)
	fmt.Println(res)
	// 1 匿名函数 调用书写
	aaa := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	println(aaa)
	// 2
	aa := func(n1, n2 int) int {
		return n1 - n2
	}
	res2 := aa(10, 20)
	println(res2)

	// 累加器   f 定义了一次， 初始化了一次变量n 引用了一次
	f := AddUpper()
	println(f(1)) // 11
	println(f(1)) // 12
}
