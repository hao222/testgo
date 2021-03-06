package main

// 将1.go 归属到main
import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// go文件 需要包  自动垃圾回收
// goroutine 轻量级线程， 基于cps并发模型实现 大并发处理
// 管道通信机制， 实现不同的goroutine相互通信
// 函数返回多个值
// 切片 延时执行defer
const test = 111

func testPtr(num *int) {
	*num = 20
}

func getSumSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sub, sum
}

func Amian() {

	c := 12
	n_1 := 5.12e2 // 5.12 * 10的2次方
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("hello", c, n_1, c1, c2)

	fmt.Printf("%T, %d, %c, %c", c, unsafe.Sizeof(c), c1, c2)

	str2 := "ab \n aaaa"
	str3 := `sssssxxxx " +
		"sadasdsada" +
		"sadasdas`
	fmt.Println("ssstr:", str2, str3)
	n1 := 111
	str := fmt.Sprintf("123%d", n1)

	fmt.Printf("str type %T %v", str, str)
	var build strings.Builder
	build.WriteString("123")
	build.WriteString("ade")
	build.WriteString("ade")
	build.WriteString("ade")
	println(build.String())

	str1 := strconv.FormatInt(int64(99), 10)
	println(str1)
	str4 := strconv.Itoa(121)
	println(str4)
	n, err := strconv.Atoi("12333")
	if err != nil {
		panic(err)
	}

	print("12+n = ", 12+n)

}

// 如果把"Print"理解为核心关键字，那么后面跟的后缀有"f"和"ln"以及""，着重的是输出内容的最终结果；
//如果后缀是"f", 则指定了format 如果后缀是"ln", 则有换行符
//如果把"Print"理解为核心关键字，那么前面的前缀有"F"和"S"以及""，着重的是输出内容的目标（终端）；
//如果前缀是"F", 则指定了io.Writer 如果前缀是"S", 则是输出到字符串
//
//如果把"Scan"理解为核心关键字，那么后面跟的后缀有"f"和"ln"以及""，着重的是输入内容的结果；
//如果后缀是"f", 则指定了format 如果后缀是"ln", 则有换行符
//如果把"Scan"理解为核心关键字，那么前面的前缀有"F"和"S"以及""，着重的是输入内容的来源（终端）；
//如果前缀是"F", 则指定了io.Reader 如果前缀是"S", 则是从字符串读取
