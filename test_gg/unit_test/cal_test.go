package main

import (
	"fmt"
	"testing"
)

// 测试用例文件名必须以 _test.go结尾， 测试用例函数必须以Test开头  Test+被测函数名， 一个测试用例文件中， 可以有多个测试用例函数
//
func TestAddUpper(t *testing.T) {

	res := AddUpper(10)
	if res != 45 {
		t.Fatalf("addupper10 报错， 实际值 %v", res) // 错误并打印出日志，退出程序
	}
	t.Logf("addupper 10 success") // 打印日志并输出
}
func TestHello(t *testing.T) {
	fmt.Println("testHello 被调用---")
}
