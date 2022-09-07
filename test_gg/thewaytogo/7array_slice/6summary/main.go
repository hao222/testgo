package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

func main() {
	//s 是一个字符串（本质上是一个字节数组
	s := "你好hello"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
	// 字符数量
	print(utf8.RuneCountInString(s))
	// 获取字符一部分, 切片是字节类型    Unicode 字符会占用 2 个字节，有些甚至需要 3 个或者 4 个字节来进行表示
	substr := s[:6]
	fmt.Println("substr:...", substr)
	// 修改字符串  通过切片来操作
	ss := "hello"
	cc := []byte(ss)
	cc[0] = 'A'
	s2 := string(cc)
	fmt.Println("s2:....", s2)
	// sort 包来实现常见的搜索和排序操作
	ints := []int{4, 7, 1, 9, 8}
	sort.Ints(ints)
	// 搜索 7 是否在里面  返回index
	xx := sort.SearchInts(ints, 7)
	fmt.Println(ints, xx)
	// append 操作
	a := []int{1, 2, 3}
	a = append(a, 88, 99, 000)
	a = append(a[:1], a[2:]...)
	// 切除 i-j 元素
	a = append(a[:1], a[3:]...)
	// 插入
	a = append(a[:1], append([]int{66}, a[1:]...)...)
	//a = append(a[:1], append(make([]int, 2), a[1:]...)...)
	b := []int{222, 333}
	a = append(a[:1], append(b, a[1:]...)...)
	var x int
	x, a = a[len(a)-1], a[:len(a)-1]
	fmt.Println("aa: ", a, x)

	// 反转字符串 两个切片实现
	str := "Google"
	sl := []byte(str)
	var rev [100]byte
	j := 0
	for i := len(sl) - 1; i >= 0; i-- {
		rev[j] = sl[i]
		j++
	}
	str_rev := string(rev[:])
	fmt.Printf("The reversed string is -%s-\n", str_rev)
	str_rev2 := reverse(str)
	fmt.Printf("The reversed string 2 is -%s-\n", str_rev2)

	list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	mf := func(i int) int {
		return i * 10
	}
	println()
	fmt.Printf("%v", mapFunc(mf, list))
}

func reverse(s string) string {
	// 取中间值反转 二分字符串
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// 冒泡排序
func bubbleSort(sl []int) {
	// passes through the slice:
	for pass := 1; pass < len(sl); pass++ {
		// one pass:
		for i := 0; i < len(sl)-pass; i++ { // the bigger value 'bubbles up' to the last position
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}
}

// 实现map函数
func mapFunc(mf func(int) int, list []int) []int {
	result := make([]int, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}
