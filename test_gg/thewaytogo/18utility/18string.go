package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 1 修改字符串的字符：
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)
	println(s2)
	// 2 获取字符串的子串
	substr := str[1:3]
	println(substr)
	// 3 使用for 是byte字节 for-range 是Unicode characters 遍历一个字符串,其实质都是字节， 需要转化
	for ix, ch := range str {
		fmt.Println(ix, ch)
		fmt.Println(ix, string(ch))
		fmt.Printf("%q", ch)
	}
	for i := 0; i < len(str); i++ {
		xx := str[i]
		fmt.Println(xx)
		fmt.Printf("%q", xx)
	}
	// 4
	fmt.Println("获取一个字符串的字节数:", len(str), "字符串的字符数", utf8.RuneCountInString(str), "or:", len([]rune(str)))
	// 字符串拼接 Strings.Join()  bytes.Buffer +=
	// os flag 命令后参数

}
