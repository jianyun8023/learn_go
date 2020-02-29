package main

import (
	"flag"
	"learn_go/demo5/lib"
	//"learn_go/demo5/lib"
	//"learn_go/demo5/lib"
	//_ "learn_go/demo5/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "Test name")

}

func main() {
	flag.Parse()
	lib.Hello(name)
}
