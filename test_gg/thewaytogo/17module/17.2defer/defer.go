package main

import "log"

// 使用 defer 可以确保资源不再需要时，都会被恰当地关闭或归还到“池子”中。更重要的一点是，它可以恢复 panic。
1 关闭文件流
defer f.Close()
2 解锁一个被锁定的资源 mutex
mu.Lock()
defer mu.Unlock()
3 关闭一个通道 如有必要
ch := make(chan float64)
defer close(ch)
4 从 panic 恢复
defer func () {
	if err := recover(); err != nil {
		log.Printf("run time panic: %v", err)
	}
}()
5
