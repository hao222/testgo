package main

import (
	"flag"
	"fmt"
	"os"
)

// os.Args  获取命令行参数 切片获取参数
// flags 包 用来解析命令行参数， 参数顺序可以随意
func main() {
	fmt.Println("---------命令行参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args %v = %v \n", i, v)
	}

	// 定义变量， 用于接收命令行参数值
	var user string
	flag.StringVar(&user, "u", "", "user name black")
	fmt.Printf("user=%v", user)
}
