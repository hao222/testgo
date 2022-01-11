package main

import (
	"fmt"
	"sync"
	"time"
)

// 阶乘 使用goroutine完成， 将各个数的阶乘放入到map中
// 出现的问题  并行并发的安全问题 ( 资源竞争问题)， 不同goroutine之间如何通信   go build -race  编译加race参数， 执行.exe的时候会发现是否资源竞争问题
var myMap = make(map[int]int, 10)
var lock sync.Mutex // 全局的互斥锁   sync 同步

func test(n int) {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()

}
func Main() {

	for i := 1; i <= 50; i++ {
		go test(i)
	}
	//// 报错  同时写入 存在资源竞争
	//for i, v := range myMap {
	//	fmt.Printf("map[%d]=%d\n", i, v)
	//}
	// 主线程休眠10s 是为了让协程完成余下的任务  其中有0 的原因是阶乘数较大，存不了
	time.Sleep(3 * time.Second)

	// 解决方案  加锁机制（全局变量加锁互斥锁   使用channel
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
	//	加锁同步实现 第二个问题 设置sleep的时间 无法预估  写 和 读 都需要加锁   所以需要channel

}
