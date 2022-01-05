package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 实现对hero结构体切片的排序， sort.Sort(data interface)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// 实现三个方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是决定使用什么标准排序
// 按age 从小到大排序  i 在 j 前
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

// 交换
func (hs HeroSlice) Swap(i, j int) {
	//tmp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = tmp
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	fmt.Println("----")
	var intslice = []int{0, -1, 10, 7, 90}
	// 排序
	sort.Ints(intslice)
	fmt.Println(intslice)
	// 对结构体排序 Sort  需要接口实现三个方法 len less swap
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄--%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	// 调用Sort方法
	fmt.Println("----排序后----")
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}

}
