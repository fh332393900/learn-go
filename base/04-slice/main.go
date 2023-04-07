package main

import "fmt"

func main() {
	// 声明切片
	var s1 []int

	fmt.Println(s1)

	s2 := []int{}
	fmt.Println(s2)

	arr := [...]int{1, 2, 3, 4, 5, 6}
	// var slice0 []int = arr[start:end]
	// 前包后不包
	s3 := arr[1:3]
	s4 := arr[:4]
	s5 := arr[3:]
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	// 通过make来创建切片
	// var slice []type = make([]type, len)

	slice := make([]int, 5, 8)
	fmt.Println(slice)

	data := [...]int{10, 20, 30, 40, 50}
	s := data[:2:4]
	s[0] += 100
	s[1] += 5
	s = append(s, 1)
	// 超过长度后
	s = append(s, 2)

	fmt.Println(data)
	fmt.Println(s)
}
