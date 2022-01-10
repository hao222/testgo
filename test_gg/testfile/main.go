package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	// 思路： 打开一个文件， 创建一个Reader
	// 每读取一行， 就去统计该行有多少个 英文 数字 空格和其他字符
	// 然后将结果保存到一个结构体

	filePath := "E:\\2021\\testgo\\test_gg\\read_file\\1.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file err=%v \n", err)
		return
	}
	defer file.Close()
	var count CharCount
	// 创建 reader
	reader := bufio.NewReader(file)
	// 循环读取file的内容
	for {
		str, err := reader.ReadString('\n') // []byte 切片
		if err == io.EOF {
			break
		}
		for _, v := range str {
			// switch v 会报错  直接使用switch {} 当做 if else
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("char %v , num %v, space %v, other %v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
