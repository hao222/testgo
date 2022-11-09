package main

import (
	"fmt"
	"sync"
)

// 新旧模型处理对比 任务和worker
// 1 旧模式： 使用共享内存进行同步  加锁机制sync.Mutex
// 加锁实现同步的方式在工作协程比较少时可以工作得很好，但是当工作协程数量很大，任务量也很多时，处理效率将会因为频繁的加锁/解锁开销而降低。
// 当工作协程数增加到一个阈值时，程序效率会急剧下降，这就成为了瓶颈
// 2 新模式： 使用通道
// 使用一个通道接受需要处理的任务，一个通道接受处理完成的任务（及其结果

func main() {
	fmt.Println("------")

}

type Task struct {
	// some state
}

type Pool struct {
	Mu    sync.Mutex
	Tasks []*Task
}

func Worker_lock(pool *Pool) {
	for {
		pool.Mu.Lock()
		// begin critical section:
		task := pool.Tasks[0]       // take the first task
		pool.Tasks = pool.Tasks[1:] // update the pool of tasks
		// end critical section
		pool.Mu.Unlock()
		process(task)
	}
}

// worker 数量的增多也会增加通信的开销，这会对性能有轻微的影响
// 如果系统部署在多台机器上，各个机器上执行 Worker 协程，Master 和 Worker 之间使用 netchan 或者 RPC 进行通信
func Worker_channel(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}
