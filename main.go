package main


import (

	"BitcoinWeb/BTCUtils/db_mysql"
	_ "BitcoinWeb/routers"
<<<<<<< HEAD
	_ "BitcoinWeb/routers"
	_ "BitcoinWeb/routers"
	"fmt"
=======
>>>>>>> 9681b73fd3f4b5e71da7a7a6f430de2514f8fe62
	"github.com/astaxie/beego"
)

func main() {

	db_mysql.ConnectDB()
    //分包处理

	beego.SetStaticPath("/js","./static/js/tool")
	beego.SetStaticPath("/jsc","./static/js")
<<<<<<< HEAD
=======
	beego.SetStaticPath("/js","./static/js/tool/")
	beego.SetStaticPath("/jsc","./static/js/")


>>>>>>> 9681b73fd3f4b5e71da7a7a6f430de2514f8fe62
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/CSS","./static/css")
	beego.SetStaticPath("/css","./static/css/Entity")

    beego.Run()
}










