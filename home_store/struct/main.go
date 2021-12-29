package main

import (
	"encoding/json"
	"fmt"
)

// struct 进行type 重新定义， go认为是新的数据类型， 但是相互之间可以强转
type Stu struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
type S Stu

// struct 的每个字段上， 可以写上一个tag``. 该tag可以通过反射机制获取，  常用于json序列化 反序列化

func main() {
	var s1 Stu
	var s2 S
	s1 = Stu{"xxx", "12"}
	//s2 = s1     需要强转
	s2 = S(s1)
	res, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("-------", s2, s1, string(res))
}
