package main

// 通道可以显示的关闭， 只有发送者需要关闭通道，
//
import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

// 使用for range 自动检测通道是否关闭
// 通道迭代器中，两个协程经常是一个阻塞另外一个。如果程序工作在多核心的机器上，大部分时间只用到了一个处理器。可以通过使用带缓冲（缓冲空间大于 0）的通道来改善。
// 比如，缓冲大小为 100，迭代器在阻塞之前，至少可以从容器获得 100 个元素。如果消费者协程在独立的内核运行，就有可能让协程不会出现阻塞
// 通道的容量需要限制一下最大值., 保证内存占用最优
func getData(ch chan string) {

	for input := range ch {
		fmt.Printf("%s ", input)
	}
	//for {
	//	input, open := <-ch
	//	if !open {
	//		break
	//	}
	//	fmt.Printf("%s ", input)
	//}
}
