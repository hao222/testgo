package main

import "fmt"

func main() {
	//arr2Dim := [...]int{0, 1, 2, 3, 43, 44}
	var arr2Dim = [4][2]int{
		{1, 0},  // 下
		{-1, 0}, // 上
		{0, 1},  // 右
		{0, -1}, // 左
	}
	found := false
Found:
	for row := range arr2Dim {
		for column := range arr2Dim[row] {
			if arr2Dim[row][column] == 1 {
				found = true
				break Found
			}
		}
	}
	arr2DimSlice := arr2Dim[:len(arr2Dim)-1]
	println(found, arr2DimSlice)
	fmt.Printf("%v", arr2DimSlice)

	// map
	map1 := map[string]int{"one": 1, "two": 2}
	for key, value := range map1 {
		println(key, value)
	}
	val1, isPresent := map1["one"] // 检测key是否存在
	println(val1, isPresent)
	delete(map1, "one") // 删除key

	// struct  当结构体的命名以大写字母开头时，该结构体在包外可见
	// 通常情况下，为每个结构体定义一个构建函数，并推荐使用构建函数初始化结构体

	//ms := new(struct1)
	//ms := &struct1{10, 15.5, "Chris"}
	//ms := Newstruct1{10, 15.5, "Chris"}
	//println(ms)

}

type struct1 struct {
	field1 int
	field2 float32
	field3 string
}

func Newstruct1(n int, f float32, name string) *struct1 {
	return &struct1{n, f, name}
}
