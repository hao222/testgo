package main

import (
	"fmt"
	"sort"
)

// 如果使用排序的列表 使用结构体切片
type name struct {
	key   string
	value int
}

//var map1 map[keytype]valuetype

// map 是可以动态增长的   未初始化值nil
// value 可以是任意类型的；通过使用空接口类型，可以储存任意值
// map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据
// 通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍

func main() {
	mp1 := make(map[int][]int)
	mp1[2] = []int{11, 22, 33}
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapCreated["key3"] = 5
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	// %v 自动匹配格式输出
	fmt.Printf("Map mp1 at \"1\" is: %v\n", mp1[2])
	delete(mp1, 2)
	_, ok := mp1[2]
	println(ok)

	// for range  map输出的结果总是无序的 因为是散列表 每次扩容都会进行重新排列
	for k, v := range mapCreated {
		fmt.Printf("key is: %v - value is: %f\n", k, v)
	}
	// map 类型的切片
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
	for k, v := range items {
		fmt.Printf("array map is: %v - value map is: %v\n", k, v)
	}

	// map的排序 默认是无序的，需要用切片 sort进行排序 切片的是key的列表
	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v ", k, barVal[k])
	}

}
