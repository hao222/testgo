package main

import (
	"fmt"
	"time"
)

// time.Tick() 函数声明为 Tick(d Duration) <-chan Time，当你想返回一个通道而不必关闭它的时候这个函数非常有用：它以 d 为周期给返回的通道发送时间，d 是纳秒数
// After它类似 Tick()，但是 After() 只发送一次时间
func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

// 超时模式  通道接收或者发送 超时
//timeout := make(chan bool, 1)
//go func() {
//	time.Sleep(1e9) // one second
//	timeout <- true
//}()
// 再使用select 接收ch or timeout数据
//select {
//case <-ch:
//// a read from ch has occured
//case <-timeout:
//// the read from ch has timed out
//break
//}
