// 从控制台读取输入:
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var (
//
//	firstName, lastName, s string
//	i                      int
//	f                      float32
//	input                  = "56.12 / 5212 / Go"
//	format                 = "%f / %d / %s"
//
// )
//var inputReader *bufio.Reader
//var input string
//var err error

var nrchars, nrwords, nrlines int

// bufio 包提供的缓冲读取器 (buffered reader) 来读取数据
// Sscan... 和以 Sscan... 开头的函数则是从字符串读取
func main() {
	//fmt.Println("Please enter your full name: ")
	//fmt.Scanln(&firstName, &lastName)
	//// fmt.Scanf("%s %s", &firstName, &lastName)
	//fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	//fmt.Sscanf(input, format, &f, &i, &s)
	//fmt.Println("From the string we read: ", f, i, s)

	//inputReader = bufio.NewReader(os.Stdin)
	//fmt.Println("Please enter some input: ")
	//input, err = inputReader.ReadString('\n')
	//if err == nil {
	//	fmt.Printf("The input was: %s\n", input)
	//}

	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "S\r\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}
}
func Counters(input string) {
	nrchars += len(input) - 2 // -2 for \r\n
	nrwords += len(strings.Fields(input))
	nrlines++
}
