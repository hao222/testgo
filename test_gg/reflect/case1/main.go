package main

import (
	"fmt"
	"reflect"
)

// 使用反射遍历结构退字段 调用结构体方法 获取结构体标签的值
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("start----------")
}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("field %d 值 %v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field %d : tag %v \n", i, tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	// 调用方法不是根据下标顺序来， 而是根据方法名的首字母ascii排序
	val.Method(1).Call(nil)
	// 参数使用切片 []reflect.Value
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	// 返回的也是切片 []reflect.Value
	fmt.Println("res=", res[0].Int())
}
func main() {

	var a Monster = Monster{
		Name:  "xxxx",
		Age:   40,
		Score: 30.1,
	}
	TestStruct(a)
	fmt.Println("-----------------")
}
