package main

import "runtime"

//DoAll() 函数创建了一个 sem 通道，每个并行计算都将在对其发送完成信号；在一个 for 循环中 NCPU 个协程被启动了，每个协程会承担 1/NCPU 的工作量。
//每一个 DoPart() 协程都会向 sem 通道发送完成信号

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

func main() {
	runtime.GOMAXPROCS(NCPU) // runtime.GOMAXPROCS = NCPU
	DoAll()
}

// 并行化大量数据的计算

// 1 典型的顺序步骤  强依赖上一个步骤的结果
func SerialProcessData(in <-chan *Data, out chan<- *Data) {
	for data := range in {
		tmpA := PreprocessData(data)
		tmpB := ProcessStepA(tmpA)
		tmpC := ProcessStepB(tmpB)
		out <- PostProcessData(tmpC)
	}
}

// 2 每一个处理步骤作为一个协程独立工作
func ParallelProcessData(in <-chan *Data, out chan<- *Data) {
	// make channels:
	preOut := make(chan *Data, 100)
	stepAOut := make(chan *Data, 100)
	stepBOut := make(chan *Data, 100)
	stepCOut := make(chan *Data, 100)
	// start parallel computations:
	go PreprocessData(in, preOut)
	go ProcessStepA(preOut, StepAOut)
	go ProcessStepB(StepAOut, StepBOut)
	go ProcessStepC(StepBOut, StepCOut)
	go PostProcessData(StepCOut, out)
}

// 漏桶算法
// c/s 结构 客户端协程执行一个无限循环从某个源头（也许是网络）接收数据；数据读取到 Buffer 类型的缓冲区
// 这个可重用的缓冲区队列 (freeList) 与服务器是共享的
