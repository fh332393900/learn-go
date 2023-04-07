package main

import (
	"fmt"
	config "learn-go/config"
)

func main() {
	fmt.Println(config.App)
	fmt.Printf("config.Mysql=%#v\n", config.Mysql)
}
