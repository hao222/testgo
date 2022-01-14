package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client is err ", err)
		return
	}
	fmt.Println("conn success", conn)
	// 循环发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read err", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.write err", err)
			return
		}
		fmt.Println("发送 字节", n)
	}

}
