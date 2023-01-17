package main

import (
	"runtime"
	"time"
)

/*
为了使并行运算获得高于串行运算的效率，在协程内部完成的工作量，必须远远高于协程的创建和相互来回通信的开销
出于性能考虑建议使用带缓存的通道：
使用带缓存的通道可以很轻易成倍提高它的吞吐量，某些场景其性能可以提高至 10 倍甚至更多。通过调整通道的容量，甚至可以尝试着更进一步的优化其性能。
限制一个通道的数据数量并将它们封装成一个数组：
如果使用通道传递大量单独的数据，那么通道将变成性能瓶颈。然而，将数据块打包封装成数组，在接收端解压数据时，性能可以提高至 10 倍。
*/

func main() {
	// 遍历通道
	//for       for-range 可以自检通道是否关闭
	// 如何通过一个通道让主程序等待直到协程完成（信号量模式）
	ch := make(chan int) // Allocate a channel.
	// Start something in a goroutine; when it completes, signal on the channel.
	go func() {
		// doSomething
		// 如果希望程序一直阻塞，在匿名函数中省略 ch <- 1 即可
		ch <- 1 // Send a signal; value does not matter.
	}()
	doSomethingElseForAWhile()
	<-ch // Wait for goroutine to finish; discard sent value.

	runtime.GOMAXPROCS(NCPU) // runtime.GOMAXPROCS = NCPU
	DoAll()
}

// 通道的工厂模板：以下函数是一个通道工厂，启动一个匿名函数作为协程以生产通道
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 如何限制并发处理请求的数量： 使用带缓冲区的通道很容易实现这一点，其缓冲区容量就是同时处理请求的最大数量
func DoAll() {
	sem := make(chan int, NCPU) // Buffering optional but sensible
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}
	// Drain the channel sem, waiting for NCPU tasks to complete
	for i := 0; i < NCPU; i++ {
		<-sem // wait for one task to complete
	}
	// All done.
}

func DoPart(sem chan int) {
	// do the part of the computation
	sem <- 1 // signal that this piece is done
}

// 如何终止一个协程   runtime.Goexit()
// 超时模板
timeout := make(chan bool, 1)
go func () {
	time.Sleep(1e9) // one second
	timeout <- true
}()
select {
case <-ch:
// a read from ch has occurred
case <-timeout:
// the read from ch has timed out
}

// 如何使用输入通道和输出通道代替锁
func Worker(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}

//从不同的并发执行的协程中获取值可以通过关键字 select 来完成，它和 switch 控制语句非常相似

// 制作， 解析并使模板生效：var strTempl = template.Must(template.New("TName").Parse(strTemplateHTML))
// 程序出错时终止程序：
if err != nil {
fmt.Printf("Program stopping with error %v", err)
os.Exit(1)
}
// or
if err != nil {
panic("ERROR occurred: " + err.Error())
}
