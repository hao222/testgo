package module

import "strings"

func Cal(n1 float64) float64 {
	var res float64
	res = n1 + 1
	return res
}

// package 区分函数 变量等 很好管理项目 控制函数 变量等访问范围， 作用域  包名通常与文件夹同名 小写
// 包名 别名  import m "module"
// 编译成一个可执行程序文件， 就需要将这个包声明为 main  即package main
// 函数调用机制， 在调用函数时， 会给该函数分配一个新的空间， 编译器会通过自身的处理让这个新的空间与其他的栈的空间区分开
// 在每个函数对应的栈中， 数据空间是独立的， 不会混淆，
// 当一个函数调用完毕后， 程序会销毁这个函数对应的栈空间

// return 返回可以返回值列表
func SumSub(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

// 函数递归调用， 终止条件为终， 分析从外向内 执行从内向外， 入栈出栈先进后出操作， 比较耗费内存
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2) // 112358
	}
}

// 函数形参列表和返回值列表可以为多个  数据类型可以为值类型和引用类型
// 首字母大写 相当于public   小写 为 private
// 基本数据类型和数组 默认值传递， 即进行值拷贝， 在函数内修改，不会影响原来的值    python 里的列表是引用会修改
// 如果让函数内的变量能修改函数外的变量， 可以传入变量地址& 函数内以*指针的方式操作变量
// 不支持重载
// 函数也是一种数据类型， 可以赋值给一个变量， 函数也可以作为形参调用
// 支持自定义数据类型 type    type myint int
// 支持函数返回值命名

type myFunc func(int, int) int

// 闭包 ： 返回的函数和外函数的变量组合，
// 该函数接收一个文件名后缀， 如果文件名没有传后缀自动加上， 如果有后缀返回原有的文件名
func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 传统函数写法  需要接收两个参数  而闭包只需要一个参数
func MakeSuffix2(name, suffix string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

// 函数中， 经常需要创建资源（数据库连接， 文件，锁）with语句...  为了在函数执行完毕后，及时的释放资源， 提供了defer（延时机制）
// 执行到defer 暂时不执行， 会将defer后面的语句压入到独立的栈defer栈
// 当函数执行完毕后， 再从defer栈， 按照先入后出的方式出栈， 执行
// 在defer将语句放入到栈时， 也会将相关的值拷贝同时入栈
// 当函数执行完毕后, 可以及时的释放函数创建的资源  defer read_file.close()
func Sums(n1, n2 int) int {
	defer println("这是defer1.。。", n1)  // 3
	defer println("这是defer2.。。。", n2) // 2
	println("defer下的----")            // 1
	n1++
	return n1 + n2
}

// 无关值传递 引用传递 传递给函数的都是变量的副本， 值传递传的是值的拷贝， 引用传递是地址的拷贝

// 全局变量 函数外部定义的， 首字母大写， 则在整个程序有效  代码块 函数内 函数外
// := 代表声明+定义+赋值  =  代表赋值或者修改值

// 字符串函数 strings
// len(str)   按照字节统计
// r := []rune(str)  字符串遍历  按照字符遍历， 同时处理有中文
// strconv.Atoi()  转整数   Itoa() 转字符串   转byte []byte("heelo")  FormatInt() 转进制
// strings.Contains()  查找子串是否在指定的字符串中  strings.Count("", "")
// 不区分大小写比较  EqualFold   返回子串第一次出现的index   strings.Index()
// strings.LastIndex  strings.Replace   Split ToLower ToUpper
// .TrimSpace() 将字符串两边的空格去掉  .Trim() 指定字符去掉 TrimLeft TrimRight

/*
日期函数 time
	time.Now()    time.Now().Unix()
	格式化必须指定时间为 2006-01-02 15:04:05
	Duration类型代表两个时间点之间经过的时间，以纳秒为单位 Second Minute Hour
	unix  将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间 秒   unixNano
*/

// 内置函数 builtin
