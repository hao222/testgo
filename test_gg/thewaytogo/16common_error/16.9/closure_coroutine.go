package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本 A：
	for ix := range values { // ix 是索引值
		func() {
			fmt.Print(ix, " ")
		}() // 调用闭包打印每个索引值
	}
	fmt.Println()
	// 版本 B：和 A 版本类似，但是通过调用闭包作为一个协程
	// 调用闭包都有一个延迟绑定问题
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本 C：调用协程的正确的处理方式   调用每个闭包时将 ix 作为参数传递给闭包
	// ix 在每次循环时都被重新赋值，并将每个协程的 ix 放置在栈中，所以当协程最终被执行时，每个索引值对协程都是可用的
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本 D：输出值：
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(1e9)
}
