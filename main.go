package main

import (
	_ "BitcoinWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Hello World!")
	beego.Run()
}

