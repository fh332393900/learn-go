package main

import "fmt"

func main() {
	var map1 = map[string]string{
		"name":    "fh",
		"address": "成都",
	}
	map2 := map[int]string{
		101: "aaa",
		300: "bbb",
	}

	fmt.Println(map1)
	fmt.Println(map2)

	// 判断键是否存在
	v, ok := map1["name"]
	fmt.Println(v, ok)

	// map 的遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// delete 删除键值对
	delete(map2, 300)
	fmt.Println(map2)
}
