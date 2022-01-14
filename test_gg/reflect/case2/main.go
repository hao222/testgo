package main

import (
	"fmt"
	"reflect"
)

// reflect 创建并操作一个结构体

type user struct {
	Userid string
	Name   string
}

func test() {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)
	st = st.Elem()
	elem = reflect.New(st)           // 返回一个value类型值， 该值持有一个指向类型为type的新申请的0值的指针
	model = elem.Interface().(*user) // model 是 *user它的指向和elem是一样的
	elem = elem.Elem()               // 取的elem指向的值
	elem.FieldByName("Userid").SetString("123")
	elem.FieldByName("Name").SetString("xxx")
	fmt.Println("model---", model.Userid, model.Name)

}
func main() {
	test()
	fmt.Println("-------------")
}
