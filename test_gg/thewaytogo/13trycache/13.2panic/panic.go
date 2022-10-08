package main

import "fmt"

// panic 程序崩溃时， 终止程序
// 当发生错误必须中止程序时，panic() 可以用于错误处理模式
// 在多层嵌套的函数调用中调用 panic()，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。
// 这样向上冒泡直到最顶层，并执行（每层） defer，在栈顶处程序崩溃，并在命令行中用传给 panic() 的值报告错误情况：这个终止过程就是 panicking

func main() {
	//fmt.Println("Starting the program")
	//panic("A severe error occurred: stopping the program!")
	//fmt.Println("Ending the program")

	// recover()) 内建函数被用于从 panic 或错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行
	// panic() 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")

}
func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- would not reach
}
