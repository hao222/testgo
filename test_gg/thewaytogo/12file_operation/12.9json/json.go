package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

//JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map[string]T（T 是 json 包中支持的任何类型）
//Channel，复杂类型和函数类型不能被编码
//不支持循环数据结构；它将引起序列化进入一个无限循环
//指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	// json.Unmarshal(js)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err1 := json.Unmarshal(b, &f)

	if err1 != nil {
		fmt.Println("error map")
	}
	fmt.Printf("unmarshal 反序列化：%v", f)
	//	json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写。
	//	NewDecoder() 和 NewEncoder() 函数分别封装了 io.Reader 和 io.Writer 接口

}
