package main

import (
	"fmt"
	"log"
	"sort"
)

type Stu struct {
	Name    string
	Age     int
	Address string
}

// map 的key 支持 bool int float string * channel interface struct array
// slice map 6function 不可以做key  无法判断
// 声明是不会分配内存的， 初始化需要make， 分配内存后才能赋值和使用  目前有序
func main() {

	// 声明后 make 再定义
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "宋江"

	aa := make(map[string]string)
	aa["n1"] = "吳1用"
	aa["n2"] = "吳用"
	delete(aa, "n2")
	log.Printf("map ---- %v ---- %v", a, aa)

	aaa := map[string]map[string]string{
		"n01": aa,
		"n02": aa,
		"n03": aa,
	}
	log.Printf("aaa ---- %v   %v", aaa, len(a))
	v, f := aaa["n1"]
	log.Printf("aaa find---- %v---%v", v, f)
	for k, v := range aaa {
		fmt.Printf("k=%v, v=%v", k, v)
		for k1, v1 := range v {
			fmt.Printf("k1=%v, v1=%v", k1, v1)
		}
	}
	println("\n")
	// slice of map
	slice_map := make([]map[string]string, 2)
	slice_map[0] = make(map[string]string, 2)
	slice_map[0]["name"] = "xxx"
	slice_map[0]["age"] = "11"
	slice_map[1] = make(map[string]string, 2)
	slice_map[1]["name"] = "xxx"
	slice_map[1]["age"] = "11"
	//// 报错 超出范围
	//slice_map[2] = make(map[string]string, 2)
	//slice_map[2]["name"] = "xxx"
	//slice_map[2]["age"] = "11"
	slice_map = append(slice_map, aa)
	log.Println("-------------", slice_map)
	// todo map 排序是根据key 升序
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 12
	map1[4] = 45
	map1[3] = 66
	fmt.Println(map1)

	//  使用sort 进行排序
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	//	 map 是引用类型  能动态增长键值对

	stu := make(map[string]Stu, 10)
	stu1 := Stu{"xx", 12, "add"}
	stu2 := Stu{"zz", 12, "add"}
	stu["no1"] = stu1
	stu["no2"] = stu2
	fmt.Println(stu)

}
