package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 接口 继承之间的关系
// 实现接口可以看做是对 继承的一种补充
// 继承的价值在于：解决代码的复用性和可维护性
// 接口的价值在于： 设计，设计好各种规范 方法， 让其它自定义类型去实现这些方法
// 接口比继承更加灵活， 继承是满足 is a的关系， 而接口只需满足like a的关系，
// 接口在一定程度上实现代码解耦

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
