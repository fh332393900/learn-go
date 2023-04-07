package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type AppConfig struct {
	// 版本号
	Version string `env:"VERSION"`
	// 运行模式
	RunMode string `env:"RUN_MODE"`
	// 服务端口
	HTTPPort int `env:"HTTP_PORT"`
	// 读写超时
	ReadTimeout  int64 `env:"READ_TIMEOUT"`
	WriteTimeout int64 `env:"WRITE_TIMEOUT"`
	// 分页每一页数量
	PageSize int `env:"PAGE_SIZE" default:"10"`
	// 默认图片
	DefaultImg string `default:"https://img0.baidu.com/it/u=3751066216,564345018&fm=253&fmt=auto&app=138&f=JPEG?w=400&h=400"`
}

var (
	App   AppConfig
	Mysql MysqlConfig
)
var config = make(map[string]string)

func init() {
	inputFile, inputError := os.Open("../.env")
	if inputError != nil {
		fmt.Printf("open env file error: %s!", inputError.Error())
		return
	}

	inputReader := bufio.NewReader(inputFile)
	var reg *regexp.Regexp
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		reg = regexp.MustCompile("\\s+")
		if reg.ReplaceAllString(inputString, "") == "" {
			continue
		}
		s := strings.Split(inputString, "=")
		s[1] = strings.Trim(s[1], "\r\n")
		s[1] = strings.Trim(s[1], "\"")
		config[s[0]] = s[1]
	}
	fmt.Println(config)
}
