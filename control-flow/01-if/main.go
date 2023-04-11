package main

import "fmt"

func main() {
	// • 可省略条件表达式括号。
	// • 持初始化语句，可定义代码块局部变量。
	// • 代码块左 括号必须在条件表达式尾部。

	// if 布尔表达式 {
	// /* 在布尔表达式为 true 时执行 */
	// }

	a := 10

	if a > 10 {
		fmt.Println("a 大于10")
	} else {
		fmt.Println("a 小于10")
	}

	b := 15

	if b > 10 && b < 20 {
		fmt.Println("b 在10-20之间")
	} else {
		fmt.Println("b 不在10-20之间")
	}

	c := 89

	if c >= 90 {
		fmt.Println("优秀")
	} else if c >= 80 {
		fmt.Println("良好")
	} else if c >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
