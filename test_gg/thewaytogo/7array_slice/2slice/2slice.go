package main

import (
	"fmt"
)

// 切片slice  长度可变的数组 引用类型
// s 是一个切片，cap(s) 就是从 s[0] 到数组末尾的数组长度
// 切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率
// 切片不能被重新分片以获取数组的前一个元素
// make 创建引用类型  适用于切片、map 和 channels 返回一个类型为 T 的初始值
// new() 函数分配内存，make() 函数初始化
//	0 <= len(s) <= cap(s)

//1.slice、map 以及 channel 都是 golang 内建的一种引用类型，三者在内存中存在多个组成部分， 需要对内存组成部分初始化后才能使用，而 make 就是对三者进行初始化的一种操作方式
//
//2. new 获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行相应的初始化操作， 所以 slice、map、channel 需要 make 进行初始化并获取对应的内存地址，而非 new 简单的获取内存地址

func main() {
	var slice1 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// 多维切片  内层的切片必须单独分配（通过 make() 函数
	// buffer  bytes
	// 切片最大值
	sl1 := []int{78, 34, 643, 12, 90, 492, 13, 2}
	max := maxSlice(sl1)
	fmt.Printf("The maximum is %d\n", max)

	// 切片创建的时候通常比相关数组小， 所以可以改变切片长度的过程  切片重组
	slice11 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice11); i++ {
		slice1 = slice11[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

func maxSlice(sl []int) (max int) {
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return
}
