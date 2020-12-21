package routers

import (
	"BitcoinWeb/controllers"
	"github.com/astaxie/beego"
)


func init() {
    beego.Router("/Bitcoin", &controllers.MainController{})
}
