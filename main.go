package main

import (
	_ "BitcoinWeb/routers"
	"github.com/astaxie/beego"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {
	fmt.Println("Hello World!")
	beego.Run()
}

