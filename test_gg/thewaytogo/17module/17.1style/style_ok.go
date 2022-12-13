package main

import "fmt"

func main() {
	// 一些事例
	// 1 在函数返回时监测错误
	value, err := pack1.Func1(param1)
	if err != nil {
		fmt.Printf("Error %s in pack1.Func1 with parameter %v", err.Error(), param1)
		return err
	}
	// 要实现简洁的错误检测代码，更好的方式是使用闭包 16.10

	// 2 检测映射中是否存在一个键值
	if value, isPresent = map1[key1]; isPresent {
		Process(value)
	}
	// 3 检测一个接口类型变量 varI 是否包含了类型 T：类型断言
	if value, ok := varI.(T); ok {
		Process(value)
	}
	//	4 检测一个通道 ch 是否关闭
	for input := range ch {
		Process(input)
	}
	// or
	for {
		if input, open := <-ch; !open {
			break // 通道是关闭的
		}
		Process(input)
	}
}
