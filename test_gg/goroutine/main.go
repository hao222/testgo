package main

import (
	"fmt"
	"strconv"
	"time"
)

// goroutine  多个协程
//如果面对随时随地可能发生的并发和线程处理需求，线程池就不是非常直观和方便了。能否有一种机制：使用者分配足够多的任务，
//系统能自动帮助使用者把任务分配到 CPU 上，让这些任务尽量并发运作。这种机制在 Go语言中被称为 goroutine。
// 并发多个任务在交替执行  并行多个任务在同一执行

func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("hello. world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	// 默认顺序执行 不使用goroutine 消耗18s  使用后 消耗9s， 实现了并发编程
	// 以主线程为主
	// 1 主线程退出， 则协程即使还没有执行完毕，也会退出，  2 协程也可以在主线程没有退出前， 就自己结束， 比如完成自己的任务
	start := time.Now()
	go test()

	for i := 1; i < 6; i++ {
		fmt.Println("hello. gogo" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	end := time.Since(start)
	fmt.Println("------", end)

}
