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
