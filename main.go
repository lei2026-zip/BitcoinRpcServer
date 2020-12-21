package main

import (
	"BitcoinRpcServer/db_mysql"
	_ "BitcoinRpcServer/routers"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("team-A's Project 2020年12月20日18:37:55")

	db_mysql.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
}

