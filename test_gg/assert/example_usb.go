package main

import "fmt"

// 类型断言用处

// 1 给Phone 结构体增加一个特有的方法call(), 当Usb接口接收的是Phone变量  还需要调用call方法

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}
type Camera struct {
}
type Computer struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机结束工作...")
}

func (p Phone) Call() {
	fmt.Println("手机 call phone--------")
}

func (c Camera) Start() {
	fmt.Println("Camera开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("Camera结束工作...")
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	// usb 为phone， 则还需要调用call方法  使用类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

	usb.Stop()
}

// 2 编写一个函数 判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("%v 的 bool类型， 值为 %v", i, x)
			fmt.Println()
		case int, int32, int64:
			fmt.Printf("%v 的 整数类型， 值为 %v", i, x)
		}
	}

}

func main() {
	fmt.Println("---------------")
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	fmt.Println("---------------")
	computer.Working(camera)
	// 一个自定义类型可以实现多个接口
	n1 := 1.11
	n2 := 20
	n3 := true
	TypeJudge(n1, n2, n3)

}
