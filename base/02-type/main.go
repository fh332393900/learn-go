package main

import "fmt"

func main() {
	const a int8 = 127
	const b uint8 = 255

	const c int32 = 10000

	const d float32 = -121.321312

	// 复数
	const e complex64 = 21 + 12i
	// bool

	// 注意：

	// 布尔类型变量的默认值为false。

	// Go 语言中不允许将整型强制转换为布尔型.

	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
	const f bool = false

	s1 := "hello"
	s2 := `第一行
    第二行
    第三行
    `
	fmt.Println(s1)
	fmt.Println(s2)

	n1 := 10
	n2 := 20
	// 类型转换
	n3 := int(float64(n1 + n2))
	fmt.Println(n3)
}
