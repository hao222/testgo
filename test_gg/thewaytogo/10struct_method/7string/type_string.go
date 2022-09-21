package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

// 如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
// 还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法
// 不要在 String() 方法里面调用涉及 String() 方法的方法
func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
