package main

import (
<<<<<<< HEAD
	"BitcoinWeb/controllers/ReadFile"
	_ "BitcoinWeb/routers"
=======
>>>>>>> 67ca010b3426f2a52ffe97ccf544d681cda094b3
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
<<<<<<< HEAD
	//这里用的绝对路径，记得修改
	fmt.Println(string(ReadFile.CommandFileReader()))
    fmt.Println("team-A's Project 2020年12月20日18:37:55")
=======
    fmt.Println("team-A's Project 2020年12月20日18:37:55-")

//	db_mysql.ConnectDB()
    //分包处理
	beego.SetStaticPath("/js","./static/js/tool/")
	beego.SetStaticPath("/jsc","./static/js/")
>>>>>>> 67ca010b3426f2a52ffe97ccf544d681cda094b3

	beego.SetStaticPath("/img","./static/img")

	beego.SetStaticPath("/CSS","./static/css")
	beego.SetStaticPath("/css","./static/css/Entity")

	beego.Run()

}

