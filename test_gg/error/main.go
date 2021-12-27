package main

import (
	"errors"
	"fmt"
)

// 错误机制  panic 错误退出程序
// 使用defer panic recover
// 抛出panic的异常， 然后再defer中通过recover捕获这个异常， 然后正常处理
func test() {

	// defer+recover 捕获 + 处理异常
	// 捕获异常 不会报错但是会终止下面的程序
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err=", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	println("-----------------")
	println(res)

}

// 自定义错误  errors.new  panic 输出错误信息， 退出程序
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件有误....")
	}
}
func test2() {
	err := readConf("confg.ini")
	if err != nil {
		panic(err)
	}
	println("test2 继续执行。。。。")
}

func main() {
	test()
	//test2()
}
