package main

import "fmt"

// channel 是一个队列 FIFO 多goroutine访问 不需要加锁， 本身线程安全  channel是有类型的， 一个string的channel只能存放string类型数据
// channel 引用类型 必须初始化make才能使用

// channel通过close关闭，关闭后只能读取数据， 而不能写数据

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	// intchan的value=0xc00011e080 intchan本身的地址=0xc000006028
	fmt.Printf("intchan的value=%v intchan本身的地址=%p\n", intChan, &intChan)
	intChan <- 10 // channel中填入数据
	num := 222
	intChan <- num

	// 管道的长度和cap（容量）
	fmt.Printf("channel len %v cap %v \n", len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	intChan <- 333
	close(intChan)
	fmt.Println("num2 ", num2)
	// 没用协程下， 如果管道数据全部取出， 再取会报错 deadlock

	// for range  channel 需要close 如果没close 会deadlock错误
	for v := range intChan {
		fmt.Println("v= ", v)
	}
}
