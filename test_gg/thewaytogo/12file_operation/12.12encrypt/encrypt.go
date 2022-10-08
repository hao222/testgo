package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

// 调用 sha1.New() 创建了一个新的 hash.Hash 对象，用来计算 SHA1 校验值。Hash 类型实际上是一个接口，它实现了 io.Writer 接口
// 通过 io.WriteString 或 hasher.Write 将给定的 []byte 附加到当前的 hash.Hash 对象中
func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "1")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	//
	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)

	hasher1 := md5.New()
	bb := []byte{}
	io.WriteString(hasher1, "test")
	fmt.Printf("md5 Result: %x\n", hasher.Sum(bb))
	fmt.Printf("md5 Result: %d\n", hasher.Sum(bb))
}
