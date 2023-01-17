package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/wuhao/PycharmProjects/goproject/testgo/target.txt")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer file.Close()
	iReader := bufio.NewReader(file)
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			return // error or EOF
		}
		fmt.Print(str)
	}
}
