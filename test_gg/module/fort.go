package module

import (
	"fmt"
	"math/rand"
	"time"
)

func Fort() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	print("\nfor循环:-----", sum)

}

// switch 基于不同表达式执行不同动作， 一个case 后面的表达式可以有多个使用逗号间隔   匹配值后面不需要加break 每个case分支都是唯一的

func Switchtest() {
	var key = 1

	switch key {
	case 1:
		print("\n这是周一")
	case 2:
		print("\n周二")
	case 3:
		print("\n周三")

	}

}

// while do... while, 需要用其他方法实现

func While() {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello", i)
		i++
	}
}

// gold star
func Gold() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Println(" ")
			}
		}
		fmt.Print(" ")
	}
}

func Multi() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}

func Breakt() {

	// 82  如果生成一个随机数， 需要先有种子seed
	// time.Now().Unix()  1970.1.1. 到现在的时间秒数
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println(n)
		count++

		if n == 99 {
			break
		}

	}
	fmt.Println("99跑了--次", count)

	// break 可以通过标签指明要终止的是哪一层语句块
lable_1:
	for i := 0; i < 3; i++ {
		//lable2:
		for j := 0; j < 3; j++ {
			if j == 1 {
				break lable_1 // 终止lable_1 的for循环
			}
			fmt.Println("jj=", j)
		}
	}
}

func Coutinue() {
	// continue 可以通过标签指明跳过哪一层循环
	// continue 结束本次循环， 进入下次循环， 起优化代码
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		println("奇数为：", i)
	}

}
