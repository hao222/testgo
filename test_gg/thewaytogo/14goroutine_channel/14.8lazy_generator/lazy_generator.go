package main

import "fmt"

// 当生成器生成数据的过程是计算密集型且各个结果的顺序并不重要时，那么就可以将生成器放入到 go 协程实现并行化
// 但是使用大量的 go 协程的开销可能会超过带来的性能增益
var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}
func generateInteger() int {
	return <-resume
}

func Main() {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())

}
