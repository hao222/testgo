package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 写文件
// openfile
func firstWrite(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return nil
	}
	// 写入 使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hello world \n")
	}
	// 内容是先写入到缓存的， 所以需要调用Flush方法， 将缓冲的数据真正写入到文件中， 否则文件中会没有数据
	writer.Flush()
	return file

}

//  os.O_APPEND 追加
func appendWrite(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return nil
	}
	// 写入 使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("你好 世界 \n")
	}
	// 内容是先写入到缓存的， 所以需要调用Flush方法， 将缓冲的数据真正写入到文件中， 否则文件中会没有数据
	writer.Flush()
	return file
}

//  os.O_RDWR 读写方式打开 并追加
func ReadWrite(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return nil
	}

	read := bufio.NewReader(file)
	for {
		str, err := read.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	// 写入 使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hahahaaha \n")
	}
	// 内容是先写入到缓存的， 所以需要调用Flush方法， 将缓冲的数据真正写入到文件中， 否则文件中会没有数据
	writer.Flush()
	return file
}

// 读取一个文件内容 并写入到另一个文件中
func test() {
	file1Path := "E:\\2021\\testgo\\test_gg\\read_file\\2.txt"
	file2Path := "E:\\2021\\testgo\\test_gg\\read_file\\3.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v \n", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file err = %v \n", err)
	}
}

// os.Stat()
// 返回的错误为nil， 说明文件或文件夹存在
// 返回的错误类型使用os.IsNotExist()判断为Ture， 说明文件或文件夹不存在
// 返回的错误为其他类型， 则不确定是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 将图片copy 到另外一个目录下   io.Copy()
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err= %v \n", err)
		return
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	filePath := "E:\\2021\\testgo\\test_gg\\read_file\\2.txt"
	// os.O_TRUNC 清空文件内容  os.O_APPEND 追加文件内容
	file := firstWrite(filePath)
	defer file.Close()
	file1 := appendWrite(filePath)
	defer file1.Close()
	file2 := ReadWrite(filePath)
	defer file2.Close()
	test()

}
