package main

import (
	"fmt"
	"runtime"
)

func main() {
	// cpuNum
	cpuNum := runtime.NumCPU()
	fmt.Println("cpunum = ", cpuNum)
	// go 默认让程序运行在多个CPU上，
}
