package main

import (
	"fmt"
	"time"
)

// 通信操作符 <-
// 同一个操作符 <- 既用于发送也用于接收  ch <- int1 流向通道（发送），从通道流出（接收）int2 = <- ch
// 通信是一种同步形式：通过通道，两个协程在通信（协程会合）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。
// 无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道。
// 在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了， 使用通道缓冲，使程序更具有伸缩性
// buf := 100
// ch1 := make(chan string, buf)

// 利用通道输出结果， 也可以利用通道达到同步目的（信号量机制），通过通道发送信号告知处理已经完成（在协程中
// 信号量模式 协程通过在通道 ch 中放置一个值来处理结束的信号。main() 协程等待 <-ch 直到从中获取到值
// 并行执行for循环的每一个迭代
// for i, v := range data {
// go func (i int, v float64) {
// doSomething(i, v)
// ...
// } (i, v)
// }

// 通道类型可以用注解来表示它只发送或者只接收 即定义只接收 或者只发送的通道
var send_chan chan<- int
var recv_chan <-chan int

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)

	// 测试阻塞性
	// 默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束
	// 1 对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果 ch 中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入
	// 2 对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了
	c := make(chan int)
	go func() {
		time.Sleep(3 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)

	// go sum
	cc := make(chan int)
	go sum(12, 13, cc)
	fmt.Println("go sum函数", <-cc) // 25

	// 两个channel 一个接受 一个导入
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRange(numChan, done)
	// main() 等待两个协程完成后再结束
	<-done

}

// 生产者消费者模式
//for {
//Consume(Produce())
//}

// 通道工厂模式   利用函数来生成通道 该函数内有匿名函数被协程调用
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}

}

// todo 通道使用for循环
func suck2(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

func numEchoRange(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func sum(x, y int, c chan int) {
	c <- x + y
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}
