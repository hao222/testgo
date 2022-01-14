package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	fmt.Println("----  redis  operate  ----")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(2))
	if err != nil {
		fmt.Println("redis dial err", err)
		return
	}
	var phone = "15690722321"
	var code = "1234"
	defer conn.Close()
	_, err = conn.Do("SET", phone, code)
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	//key失效时间
	_, err = conn.Do("expire", phone, 20)
	if err != nil {
		fmt.Println("set expire error: ", err)
	}
}
