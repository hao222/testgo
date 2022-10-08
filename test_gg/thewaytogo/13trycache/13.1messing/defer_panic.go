package main

// Go 没有像 Java 和 .NET 那样的 try/catch 异常机制：不能执行抛异常操作。但是有一套 defer-panic-and-recover 机制
// 产生错误的函数会返回两个变量，一个值和一个错误码；如果后者是 nil 就是成功，非 nil 就是发生了错误。
//
// 为了防止发生错误时正在执行的函数（如果有必要的话甚至会是整个程序）被中止，在调用函数后必须检查错误
//通常你想要返回包含错误参数的更有信息量的字符串，例如：可以用 fmt.Errorf() 来实现：它和 fmt.Printf() 完全一样，
//接收一个或多个格式占位符的格式化字符串和相应数量的占位变量。和打印信息不同的是它用信息生成错误对象
import (
	"errors"
	"fmt"
)

// 定义一个error
var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
	ff, err := Sqrt(-1)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("%v", ff)

	// 对于不同的错误类型 使用type cache
	//switch err := err.(type) {
	//case os.ParseError:
	//	PrintParseError(err)
	//case PathError:
	//	PrintPathError(err)
	//	...
	//default:
	//	fmt.Printf("Not a special error, just %s\n", err)
	//}

	// 从命令行读取输入， -h --help 添加一些错误
	//if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
	//	err = fmt.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
	//	return
	//}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	// implementation of Sqrt
	return f * 0.1, nil
}
