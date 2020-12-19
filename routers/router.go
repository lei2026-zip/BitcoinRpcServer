package routers

import (
	"BitcoinWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/come_in", &controllers.MainController{})

}
