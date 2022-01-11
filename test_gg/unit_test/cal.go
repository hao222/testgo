package main

// 被测试的函数
func AddUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}
