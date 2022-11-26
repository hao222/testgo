package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"test_gg/test_gg/thewaytogo/15net_web_/15.9rpc/rpc_objects"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
