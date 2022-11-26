package main

import (
	"fmt"
	"log"
	"net/rpc"
	"test_gg/test_gg/thewaytogo/15net_web_/15.9rpc/rpc_objects"
)

const serverAddress = "localhost"

// 客户端必须知晓对象类型及其方法的定义。执行 rpc.DialHTTP() 连接到服务器后，就可以用 client.Call("Type.Method", args, &reply) 调用远程对象的方法。
// Type 是远程对象的类型名，Method 是要调用的方法，args 是用 Args 类型初始化的对象，reply 是一个必须事先声明的变量，方法调用产生的结果将存入其中
func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpc_objects.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}

// 一项和 rpc 密切相关的技术是基于网络的通道，类似 14 章所使用的通道都是本地的，它们仅存在于被执行的机器内存空间中
//netchan 包实现了类型安全的网络化通道：它允许一个通道两端出现由网络连接的不同计算机
