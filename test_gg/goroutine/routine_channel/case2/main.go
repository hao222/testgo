package main

import "fmt"

func putNum(intChan chan int) {
	for i := 0; i < 200; i++ {
		intChan <- i
	}
	// 关闭
	close(intChan)

}

// 开启4个协程， 从intChan取出数据， 判断是否为素数，放入primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { // intchan 取不到
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// 为素数
			primeChan <- num
		}
	}
	// 不要直接关闭primeChan 因为你不确定其他协程是否在用  使用向exitchan 写入true
	// exitchan 开启4个协程 如果exitchan有4个true 则表明协程都已经运行完毕， 可以退出
	exitChan <- true
}

func main() {
	intChan := make(chan int, 100)
	primeChan := make(chan int, 200)
	exitChan := make(chan bool, 4) // 标识退出的管道
	fmt.Println("-----")

	// insert intChan data
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// 若从exitChan取出4个， 可以关闭primeChan
		close(primeChan)
	}()
	// 遍历primeNum  将结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数= %d\n", res)
	}
	fmt.Println("exit main process")
}
