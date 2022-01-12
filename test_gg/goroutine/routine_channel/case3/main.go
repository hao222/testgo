package main

import "fmt"

func main() {
	// 使用select 可以解决从管道取数据的阻塞
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// 传统遍历管道 如果不关闭会阻塞 deadlock
	// 不关闭管道 使用select 读取数据
	for {
		select {
		case v := <-intChan:
			fmt.Printf("intcahn取出数据 %d\n", v)
		case v := <-stringChan:
			fmt.Printf("intcahn取出数据 %s\n", v)
		default:
			fmt.Printf("none...")
			return
		}
	}

	// goroutine 使用recover 解决协程中出现的panic。 导致程序崩溃问题

}
