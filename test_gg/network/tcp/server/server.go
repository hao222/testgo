package main

import (
	"fmt"
	"net"
)

func Process(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close()
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// 等待客户端通过conn发送信息，如果客户端没有write 则协程会阻塞在这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("receive err...", err)
			return
		}
		// 服务端接收数据
		fmt.Print("-------", string(buf[:n]))
	}
}

func main() {

	fmt.Println("server is starting...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err= ", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err = ", err)
		}
		//准备go 协程 为客户端服务
		go Process(conn)
		fmt.Println("accept connect =  conn ip", conn, conn.RemoteAddr())
	}
}
