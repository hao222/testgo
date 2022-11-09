package main

// 所谓 Futures 就是指：有时候在你使用某一个值之前需要先对其进行计算
// Futures 模式通过闭包和通道可以很容易实现，类似于生成器，不同地方在于 Futures 需要返回一个值

func InverseFuture(a Matrix) chan Matrix {
	future := make(chan Matrix)
	go func() {
		future <- Inverse(a)
	}()
	return future
}

func InverseProduct(a Matrix, b Matrix) {
	a_inv_future := InverseFuture(a) // start as a goroutine
	b_inv_future := InverseFuture(b) // start as a goroutine
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return Product(a_inv, b_inv)
}
