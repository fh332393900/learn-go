package main

import "fmt"

func main() {
	a := 1
	c := 10
	p := &a
	// 指针取值
	var b int = *p
	fmt.Println(b)
	fmt.Printf("type of p:%T\n", p)
	fmt.Printf("type of c:%T\n", b)
	fmt.Println(p)
	fmt.Println(&c)

	x, y := position(p, &c)

	fmt.Println(x, y)

	// 空指针
	var p2 *string

	if p2 != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
}

// 指针函数传值
func position(x *int, y *int) (int, int) {
	*x += 10
	*y += 10
	fmt.Println(x, "x")
	fmt.Println(y, "y")
	return *x, *y
}
