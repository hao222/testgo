package main

import (
	"fmt"
	"io/ioutil"
)

// 一个go目录只允许出现一个main文件 如果想执行多个main文件  通过 go run xxx.go

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("%s\n", contents)

	}
	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Printf("%s\n", contents)
	// }
	fmt.Println(111)
}
