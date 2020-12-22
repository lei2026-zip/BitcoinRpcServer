package main

import (
	"BitcoinWeb/controllers/ReadFile"
	_ "BitcoinWeb/routers"
	_ "BitcoinWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//这里用的绝对路径，记得修改
	fmt.Println(string(ReadFile.CommandFileReader()))

    fmt.Println("team-A's Project 2020年12月20日18:37:55-")

//	db_mysql.ConnectDB()
    //分包处理
	beego.SetStaticPath("/jsc","./static/js/")
	beego.SetStaticPath("/js","./static/js/tool")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/CSS","./static/css")
	beego.SetStaticPath("/css","./static/css/Entity")
	beego.Run()

}

