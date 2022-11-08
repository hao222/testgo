package main

import (
	"fmt"
	"time"
)

// 不要使用全局变量或者共享内存，它们会给你的代码在并发运算的时候带来危险
// 应用程序并发处理的部分被称作 goroutines（协程），它可以进行更有效的并发运算
// 使用 channels 来同步协程
// 存在两种并发方式：确定性的（明确定义排序）和非确定性的（加锁/互斥从而未定义排序）
// Go 的协程和通道理所当然的支持确定性的并发方式（例如通道具有一个 sender 和一个 receiver）
// 协程的栈会根据需要进行伸缩，不出现栈溢出
// main函数也可以看做一个协程， 协程可在初始化init()函数中运行
// 对于 n 个核心的情况设置 GOMAXPROCS 为 n-1 以获得最佳性能 , 协程的数量 > 1 + GOMAXPROCS > 1
// 当 main() 函数返回的时候，程序退出：它不会等待任何其他非 main 协程的结束。
// 这就是为什么在服务器程序中，每一个请求都会启动一个协程来处理，server() 函数必须保持运行状态。通常使用一个无限循环来达到这样的目的
// Go 协程意味着并行（或者可以以并行的方式部署），协程一般来说不是这样的
// Go 协程通过通道来通信；协程通过让出和恢复操作来通信
func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	// todo 协程会随着程序的结束而消亡
	time.Sleep(1 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(2 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
