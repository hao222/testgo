package main

import (
	"fmt"
)

// 切片的两个属性长度和容量, 可以通过内建函数len和cap获得  底层为struct
// 切片的长度是它所包含的元素个数。  切片是引用类型
// 切片的容量是从它的第一个元素到其底层数组元素末尾的个数。 cap 最大能存储多少元素
// new 只分配内存，而 make 只能用于 slice、map 和 channel 的初始化
// 切片初始化时, 仍然不能越界, 范围在 0- len(arr)  可以动态增长
// copy 拷贝切片
func slicetest() {
	// 切片是数组的一个引用
	var s2 = []int{1, 2}
	var arr = [...]int{1, 2, 3, 4}
	// 1
	slice := arr[1:3]
	// 2   默认值全为 当前类型的默认值
	slice2 := make([]int, 4, 10)
	// for len 遍历  for range 遍历
	fmt.Println("-------", s2, slice, slice2)
	var a = []int{100, 200}
	// 动态追加 append 本质是对数组扩容，将slice2原来的元素拷贝到新数组
	slice2 = append(slice2, a...)
	// 都为切片类型  后一个copy到前一个
	copy(slice2, slice)
	fmt.Println("slice2===", slice2)
	// string 切片修改， 通过将string ---> []byte --->string()  []rune -->修改--> string()
	// []byte 对中文会出现乱码  改成 []rune
	aa := "ewfeqe32"
	aa1 := []rune(aa)
	fmt.Println("aa----", string(aa1))

}

// 冒泡排序
func Bubble(arr *[5]int) {

	tmp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}
	fmt.Println("排序后===", (*arr))
}

// 二分查找
func Binary(arr *[6]int, leftIndex int, rightIndex int, findin int) {

	if leftIndex > rightIndex {
		println("==找不到")
		return
	}
	mid := (leftIndex + rightIndex) / 2
	if (*arr)[mid] > findin {
		Binary(arr, leftIndex, mid-1, findin)
	} else if (*arr)[mid] < findin {
		Binary(arr, mid+1, rightIndex, findin)
	} else {
		fmt.Printf("====下标为   %v \n", mid)
	}

}

// 二维数组
func arr2() {

	var arr [4][6]int
	arr[1][2] = 2
	arr[2][1] = 1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	var arr2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i, v := range arr2 {
		fmt.Printf("arr3 %v %v", i, v)
	}

}

// 实现在一个排序好升序的数组内插入一个数， 按照原来的排序规律
func orderList() {
	var arr = [6]int{1, 4, 8, 10, 22}
	var tmp int
	var tmp2 int
	inserta := 3
	if inserta > arr[4] {
		arr[5] = inserta
	} else {
		for i := 0; i < 6; i++ {
			if arr[i] > inserta {
				tmp = arr[i]
				arr[i] = inserta
				for j := i + 1; j < 6; j++ {
					tmp2 = arr[j]
					arr[j] = tmp
					tmp = tmp2
				}
				break
			}
		}
	}
	for i := 0; i < 6; i++ {
		fmt.Printf(" %v", arr[i])
	}
	println()

}

// &myarr 数组地址   数组第一个元素地址就是数据的首地址
// 数组是多个相同类型数据的组合， 一旦声明定义， 长度是固定的，不能动态变化 可以是任何数据类型，不能混用 创建后会有默认值
// go的数组是值类型， 默认为值传递， 数组间不会相互影响   如果想在函数中修改该数组值， 使用引用方式 在函数 数组参数传递的时候， 要考虑到长度也是数组的一部分
func main() {
	var myarr = [3]int{1, 2, 3}
	var myarr1 = [...]int{5, 6, 7}
	var myarr2 = [...]int{0: 10, 1: 100} // 指定index
	c := [2]string{"yes", "no"}
	fmt.Println(myarr, myarr1, myarr2, c, len(myarr2))
	c[1] = "yyyy"
	fmt.Println(c)
	// for range 结构遍历   for len()
	for i, v := range myarr {
		fmt.Printf("index =%v, vv = %v", i, v)
	}
	for _, v := range myarr {
		fmt.Printf("vv = %v", v)
	}

	lens := len(myarr1)
	tmp := 0
	for i := 0; i < lens/2; i++ {
		tmp = myarr1[lens-1-i]
		myarr1[lens-1-i] = myarr1[i]
		myarr1[i] = tmp
	}
	fmt.Println("反转的myarr1", myarr1)
	slicetest()
	arr1 := [5]int{4, 7, 1, 3, 7}
	Bubble(&arr1)
	arr := [6]int{4, 7, 1, 3, 7, 2}
	Binary(&arr, 0, len(arr)-1, 3)
	arr2()
	orderList()
}
