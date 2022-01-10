package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读文件 输入流 数据从数据源（文件） 到程序（内存）的路径
// 写文件 输出流 数据从程序（内存） 到数据源（文件）的路径 os.File 结构体

// 创建一个reader  带缓冲的读文件  defaultBufSize = 4096 默认缓冲区为4096
func Reader(file *os.File) {
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到换行就结束
		if err == io.EOF {                  // eof 文件的末尾
			break
		}
		// 输出
		fmt.Print(str)
	}
}

// 一次性读取文件内容并显示在终端， 适用于文件不大的情况
func ReadFile() {
	file := "E:\\2021\\testgo\\test_gg\\read_file\\1.txt"
	//文件的open 和 close 被封装到readfile函数内部
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read fiel %v", err)
	}
	fmt.Printf("%v", content) // []byte
	fmt.Printf("%s", content) // string 类型

}

func main() {
	file, err := os.Open("E:\\2021\\testgo\\test_gg\\read_file\\1.txt")
	if err != nil {
		fmt.Println("open read_file", err)
	}
	// read_file 文件对象  文件句柄 file指针
	fmt.Printf("-----%v-", file)
	//err = read_file.Close()
	//if err != nil{
	//	fmt.Println("close read_file err=", err)
	//}
	// 当函数退出时， 要及时的关闭file   要及时关闭file句柄， 否则会有内存泄露
	defer file.Close()
	Reader(file)
	ReadFile()
}
