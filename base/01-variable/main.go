package main

import "fmt"

func main() {
	// var 变量名 变量类型
	var name string
	var age int

	// 批量声明
	var (
		address string = "四川省成都市"
		email   string = "111@qq.com"
	)

	// 短变量声明
	car := "BMW"

	// 匿名变量 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
	x, _ := foo()

	const pi = 3.14
	const e = 2.71
	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const (
		n1 = 100
		n2
		n3
	)

	// iota是go语言的常量计数器，只能在常量的表达式中使用。

	const (
		a = iota // 0
		b        // 1
		c        // 2
	)
	// 使用_跳过某些值
	const (
		m1 = iota // 0
		m2        // 1
		_
		m3 // 3
	)
	// iota声明中间插队
	const (
		A = iota
		B = 666
		C = iota
		D
	)

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(address)
	fmt.Println(email)

	fmt.Println(car)

	fmt.Println("x=", x)

	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}

func foo() (int, string) {
	return 10, "fh"
}

// 函数外的每个语句都必须以关键字开始（var、const、func等）

// :=不能使用在函数外。

// _多用于占位，表示忽略值。
