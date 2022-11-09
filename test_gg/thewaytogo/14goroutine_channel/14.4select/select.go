package main

import (
	"fmt"
	"time"
)

// 使用select切换协程
// 从不同的并发执行的协程中获取值可以通过关键字 select 来完成， select case
// select 做的就是：选择处理列出的多个通信情况中的一个， 如果多个符合则会随机选择一个
// 在select中使用发送操作并且有 default 可以确保发送不被阻塞

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e6)

	// 使用通道实现斐波那契数列
	c := make(chan int, 10)
	go fib(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fib(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 10
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("received on channel1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("received on Channel 2: %d\n", v)
		}

	}
}
