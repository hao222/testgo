package main

import (
	"bufio"
	"fmt"
	"os"
)

//io/ioutil 包里的 ioutil.ReadFile() 方法，该方法第一个返回值的类型是 []byte，里面存放读取到的内容，
//第二个返回值是错误，如果没有错误发生，第二个返回值为 nil

//compress 包提供了读取压缩文件的功能，支持的压缩文件格式为：bzip2、flate、gzip、lzw 和 zlib。

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string
	// todo 读文件os.open
	//inputFile, inputError := os.Open("/Users/wuhao/mix/jx_dt/tets.txt")
	//if inputError != nil {
	//	fmt.Printf("An error occurred on opening the inputfile\n" +
	//		"Does the file exist?\n" +
	//		"Have you got access to it?\n")
	//	return // exit the function on error
	//}
	//defer inputFile.Close()
	//
	//inputReader := bufio.NewReader(inputFile)
	//
	//for {
	//	inputString, readerError := inputReader.ReadString('\n')
	//	fmt.Printf("The input was: %s", inputString)
	//	if readerError == io.EOF {
	//		return // error or EOF
	//	}
	//}

	//inputFile := "/Users/wuhao/mix/jx_dt/tsets.txt"
	//outputFile := "/Users/wuhao/mix/jx_dt/tets_copy.txt"
	//buf, err := ioutil.ReadFile(inputFile)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	//	// panic(err.Error())
	//}
	//fmt.Printf("%s\n", string(buf))
	//err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//filename := filepath.Base("/Users/wuhao/mix/jx_dt/tets.txt")
	//println("filename", filename)

	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	// 使用缓冲区
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	//在缓冲写入的最后千万不要忘了使用 Flush()，否则最后的输出不会被写入
	outputWriter.Flush()

	//  不使用缓冲区
	//os.Stdout.WriteString("hello, world\n")
	//f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	//defer f.Close()
	//f.WriteString("hello, world in a file\n")

}
