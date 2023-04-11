package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var str string = "abcdefg"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	a := 5
	for a > 0 {
		fmt.Println(a)
		a--
	}
	// 无限循环
	// for {
	// 	println(str)
	// }
}
