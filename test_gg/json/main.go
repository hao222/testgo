package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"` // 反射机制
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

// 将结构体序列化
func testStruct() {
	// 增加 tag标签
	monster := Monster{
		Name:     "xiaoming",
		Age:      122,
		Birthday: "1999-02-01",
		Sal:      80000.0,
		Skill:    "qiongchacha",
	}
	// 将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化err %v\n", err)
	}
	fmt.Printf("monster序列化后=%v", string(data))
}

func testMap() {
	// map key is string  value is interface 任意类型
	var a map[string]interface{}
	// 使用map前 需要make
	a = make(map[string]interface{})
	a["name"] = "xiaosan"
	a["age"] = 30
	a["address"] = "浦东新区"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化err %v\n", err)
	}
	fmt.Printf("a 序列化=%v", string(data))
}

// 对切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	// 使用map 先make
	m1 = make(map[string]interface{})
	m1["name"] = "lucky"
	m1["age"] = 21
	m1["address"] = "南京"
	slice = append(slice, m1)
	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "lucy"
	m2["age"] = 21
	m2["address"] = "西京"
	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化err %v\n", err)
	}
	fmt.Printf("切片 序列化=%v", string(data))
}

// 基本数据类型的格式化
func testFloat() {
	var num1 float64 = 234.21
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化err %v\n", err)
	}
	fmt.Printf("float基本类型 序列化=%v", string(data))
}
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat()
}
