package main

import "fmt"

// go中 多态的特性是通过接口实现的   高内聚 低耦合
// interface 类型可以定义一组方法， 但是这些不需要实现， 并且 interface不能包含任何变量， 在某个自定义类型使用方法的时候， 在具体将这些方法写出
// 接口里的所有方法都没有方法体， 即接口的方法都是没有实现的方法， 接口体现了程序设计的多态和高内聚低耦合的思想
// go的接口 不需要显式的实现， 只要一个变量， 含有接口类型中的所有方法， 那么这个变量就实现这个接口。 因此，go中没有implement这样的关键字
// 需要全部实现接口下的方法， 才能说该自定义类型实现了该方法
// 接口本身不能创建实例， 但是可以指向一个实现了该接口的自定义类型的变量
// var phone Phone
// var usb Usb = phone
// usb.Stop()

// 一个自定义类型可以实现多个接口
// go接口中不能有任何变量
// 一个接口可以集成多个别的接口， 这时如果要实现A接口， 也必须将BC接口方法都实现
// interface类型默认是一个指针， 如果没有对interface初始化就使用，那么会输出nil
// 空接口interface{} 没有任何方法， 所以所有类型都实现了空接口

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
func (c Camera) Start() {
	fmt.Println("Camera开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("Camera结束工作...")
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
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
	Main()

}
