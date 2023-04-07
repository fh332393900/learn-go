package main

import "fmt"

func main() {
	type person struct {
		name  string
		age   int32
		phone string
	}

	var fh person
	fh.age = 25
	fh.name = "fenghang"

	// 匿名结构体
	var car struct {
		speed int
		name  string
	}
	car.speed = 120

	fmt.Println(fh)
	fmt.Println(car)

	// 取结构体的地址实例化
	fh2 := &person{}

	fh2.name = "fenghang2"
	fmt.Println(fh2)
	fmt.Printf("p3=%#v\n", fh2)
}
