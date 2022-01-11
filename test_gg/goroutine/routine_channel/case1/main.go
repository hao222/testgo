package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("write v = %v \n", i)

	}
	close(intChan) //
}

func readData(intChan chan int, exitChan chan bool) {
	// 如果只向管道写入数据， 而没有读取， 则会发生阻塞导致deadlock
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("read v = %v \n", v)
	}
	exitChan <- true
	close(exitChan)

}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)
	fmt.Println("---------------")
	// 第一种方式 使用time.sleep
	//time.Sleep(time.Second * 5)
	// 第二种 使用channel 方式连接goroutine
	for {
		ok := <-exitChan
		if !ok {
			break
		}
	}

}
