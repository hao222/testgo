// big.go
package main

import (
	"fmt"
	"math"
	"math/big"
	"test_gg/test_gg/thewaytogo/9package/4big_jingmijisuan/self_add"
)

//import . "./pack1/pack1"  // 别名  . 可以直接使用pack1下的变量方法等  如果为：_ 则只执行包内的init函数和全局变量

func main() {
	// Here are some calculations with bigInts: 大整数
	println(self_add.PFloat)
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts: 大有理数  分子 分母
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}
