package main

import "fmt"

func main() {
	arr := [...]int{10, 20, 30, 40, 50}

	for index, value := range arr {
		fmt.Println(index, "index")
		fmt.Println(value, "value")
	}

	map1 := map[string]int{
		"a": 10,
		"b": 80,
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}
	// range 会复制对象。
}
