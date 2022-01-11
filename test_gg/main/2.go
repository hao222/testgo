package main

import (
	"fmt"
	"strconv"
	"time"
)
import "test_gg/test_gg/module"

func main() {
	str := "true"
	var bb bool
	// _ 忽略返回的值
	bb, _ = strconv.ParseBool(str)
	print(bb)
	i := 10
	var ptr *int = &i
	print(ptr, "'----'", *ptr)
	*ptr = 111
	module.Add(1, 2)
	// 去掉小数部门 保留整数
	print(10/4, "\n")
	print(10.0 / 4)
	print("取余：", -10%-3) // -1

	a := 10
	b := 20
	a = a + b
	b = a - b
	a = a - b
	print("\n", a, b)

	module.Fort()

	//module.Scanll()
	// 二进制 补码方式运行
	print("\n", 2&3, "----", 2|3, "-----", 2^3)

	// 在if判断条件中直接定义变量
	if age := 18; age >= 18 {
		print("已成年")
	} else if name := "xiao"; name == "xiao" {
		print("实小")
	} else {
		print("未成年")
	}
	var year int = 2021
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		print("闰年")
	} else {

		print("\n不是闰年")
	}

	module.Switchtest()
	module.While()

	module.Gold()
	module.Multi()
	print("\n")
	module.Breakt()
	module.Coutinue()
	// goto 语句可以无条件地转义到程序中制定的行  默认不使用
	goto label1
	println("xxxx")
label1:
	println("fffff")

	println("ggggg")
	res := module.Cal(1)
	println(res)
	sum, sub := module.SumSub(9, 2)
	println(sum, sub)
	fib := module.Fib(10)
	println(fib)

	rr := module.Sums(1, 2)
	println(rr)

	println("str--len:", len("hlellp")) //
	now := time.Now()
	fmt.Printf(now.Format("2006-01-02 15:04:05"))

	num2 := new(int)
	*num2 = 110
	fmt.Printf("类型%T 值地址%v 地址%v  指针值%v", num2, num2, &num2, *num2)
}
