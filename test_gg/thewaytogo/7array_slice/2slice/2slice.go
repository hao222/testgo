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

func maxSlice(sl []int) (max int) {
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return
}

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

	// 切片的追加复制
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3
	// 若 s 的容量不足以存储新增元素，append() 会分配新的切片来保证已有切片元素和新增元素的存储
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)

	ss := []string{"M", "N", "O", "P", "Q", "R"}
	res := RemoveStringSlice(ss, 2, 4)
	fmt.Println(res) // [M N Q R]
}

// filter方法从 slice中获取一些需要的数
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func RemoveStringSlice(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	fmt.Println("..........at", at)
	copy(result[at:], slice[end:])
	return result
}
