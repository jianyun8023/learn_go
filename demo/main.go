package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	//覆写flag的使用说明
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	// 定义参数name，默认值`everyone`
	flag.StringVar(&name, "name", "everyone", "Test name")

}

func main() {
	//格式化参数
	flag.Parse()
	fmt.Printf("Hello %s!\n", name)
}
