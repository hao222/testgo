package main

import "fmt"

// 多态是通过接口实现的
// 多态参数 Usb usb  可以接收手机变量，又可以接收相机变量， 体现了usb接口多态
// 多态数组 给Usb数组中，存放phone结构体和camera结构体变量
// var usbArr [3]Usb   usbArr[0] = Phone{"vivo"}  usbArr[1] = Camera{"NiKang"}

func main() {
	fmt.Println("-----")
}
