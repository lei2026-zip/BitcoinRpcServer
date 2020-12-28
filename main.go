package main


import (

	"BitcoinWeb/BTCUtils/db_mysql"
	_ "BitcoinWeb/routers"
	"github.com/astaxie/beego"
)

func main() {

	db_mysql.ConnectDB()
    //分包处理

	beego.SetStaticPath("/js","./static/js/tool")
	beego.SetStaticPath("/jsc","./static/js")
	beego.SetStaticPath("/js","./static/js/tool/")
	beego.SetStaticPath("/jsc","./static/js/")


	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/CSS","./static/css")
	beego.SetStaticPath("/css","./static/css/Entity")

    beego.Run()
}

