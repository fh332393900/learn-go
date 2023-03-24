package main

import "fmt"

func main() {
	/**
	数组定义：var a [len]int，比如：var a [5]int，
	数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。

	数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
	*/
	var arr1 [5]int = [5]int{1, 2, 3}
	fmt.Println(arr1[2]) // 3
	fmt.Println(arr1[4]) // 0

	// 不限长度数组
	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// 使用索引号初始化元素。
	arr3 := [...]int{2: 10, 4: 30}
	fmt.Println("arr3", arr3)

	arr4 := [...]struct {
		name string
		age  int8
	}{
		{name: "fh", age: 25},
	}
	fmt.Println("arr4", arr4)

	// 数组拷贝
	var arr5 [5]int
	var arr6 = &arr5
	arr5 = [5]int{1, 2, 3, 4, 5}
	arr6[0] = 10
	fmt.Println("arr5", arr5)
	fmt.Println("arr6", *arr6)

	// 多维数组 第 2 纬度不能用 "..."。
	arr7 := [...][3]int{{1, 3, 4}, {3, 4, 5}}
	fmt.Println(arr7)
}
