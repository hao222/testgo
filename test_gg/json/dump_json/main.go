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

// json的反序列化  Unmarshal  要确保反序列化后的数据类型和原来序列化前的数据类型一致
// 反序列化struct
func unmarshStruct() {
	str := "{\"name\":\"xiaoming\",\"age\":122,\"Birthday\":\"1999-02-01\",\"Sal\":80000,\"Skill\":\"qiongchacha\"}"
	var mon Monster
	err := json.Unmarshal([]byte(str), &mon)
	if err != nil {
		fmt.Printf("unmash err %v", err)
	}
	fmt.Printf("反序列化后 mon=%v", mon)
}

func main() {
	fmt.Println("------")
	unmarshStruct()
}
