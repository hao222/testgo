package main

import "bytes"

// todo 误用短声明导致变量覆盖，应使用 =
var remember bool = false

if something {
remember := true // 错误   将会只能在if 语句中生效， 外层的变量仍然是原来值
}

if something {
remember = true // 正确   会将外层变量修改
}

// todo go 语言中字符串是不可变的， 应使用一个字符数组代替字符串，将字符串内容写入一个缓存中
var b bytes.Buffer
...
for condition {
b.WriteString(str) // 将字符串str写入缓存buffer
}
return b.String()

// 发生错误时使用 defer 关闭一个文件
// 如果你在一个 for 循环内部处理一系列文件，你需要使用 defer 确保文件在处理完毕后被关闭，例如：
for _, file := range files {
if f, err = os.Open(file); err != nil {
return
}
// 这是错误的方式，当循环结束时文件没有关闭
defer f.Close()
// 对文件进行操作
f.Process(data)
}
// 但是在循环内结尾处的 defer 没有执行，所以文件一直没有关闭！垃圾回收机制可能会自动关闭文件
// todo defer 仅在函数返回时才会执行，在循环内的结尾或其他一些有限范围的代码内不会执行

// todo 切片、映射和通道，使用 make()  数组、结构体和所有的值类型，使用 new()

// todo 切片当作参数需要注意 切片slice底层通过数组实现，slice类似一个结构体，其中一个字段保存的是底层数组的地址，还有长度(len) 和 容量(cap）两个字段
// 注意是值传递还是引用传递

// todo 永远不要使用一个指针指向一个接口类型，因为接口类型已经是一个指针

// 在大多数情况下，通过栈传递参数会更有效率
// 但是，如果你使用 break、return 或者 panic() 去跳出一个循环，很有可能会导致内存溢出，因为协程正处理某些事情而被阻塞。
// 在实际代码中，通常仅需写一个简单的过程式循环即可。当且仅当代码中并发执行非常重要，才使用协程和通道
