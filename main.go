package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("team-A's Project 2020年12月20日18:37:55-")

//	db_mysql.ConnectDB()
    //分包处理
	beego.SetStaticPath("/js","./static/js/tool/")
	beego.SetStaticPath("/jsc","./static/js/")

	beego.SetStaticPath("/img","./static/img")

	beego.SetStaticPath("/CSS","./static/css")
	beego.SetStaticPath("/css","./static/css/Entity")

	beego.Run()
}

