package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

/** 声明并创建一个cmd解析对象 */
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	/** 声明参数 */
	cmdLine.StringVar(&name, "name", "everyone", "Test name")

}

func main() {
	/**解析参数，启动参数0位表示当前可执行程序，要排除掉 */
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello %s!\n", name)
}
