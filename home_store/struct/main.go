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

type student struct {
	Name string
	score float64
}

// 工厂模式  如果属性时小写 也不能在其他包直接获取， 可以定义一个方法获取
func Factory(n string ,s float64) *student {
	return &student{
		Name: n,
		score: s,
	}
}
// 也是一种封装  设置set get方法
func (s *student) GetScore() float64 {
	return s.score
}

// struct 的每个字段上， 可以写上一个tag``. 该tag可以通过反射机制获取，  常用于json序列化 反序列化
// 抽象  将具体的事务的属性描述到公共的字段上表示 代表一类的事物
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

	// 在创建结构体变量时， 把字段名和字段值写在一起， 这种写法， 就不依赖字段的定义顺序
	var s3 = Stu{
		Name: "xiaoming",
		Age:  "12",
	}
	s5 := &Stu{"xxxx", "34"}
	fmt.Println(*s5)
	fmt.Println("-------", s2, s1, string(res), s3)
	// 定student结构体首字母小写的时候， 让其他包能导入 除了首字母大写之外， 通过工厂模式写
	var stu = student{"toms", 22.2}
	fmt.Println(stu.Name, stu.GetScore())
}
