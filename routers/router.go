package routers

import (
	"BitcoinWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
<<<<<<< HEAD
    beego.Router("/come_in", &controllers.MainController{})

=======
    beego.Router("/", &controllers.MainController{})
>>>>>>> 63255a460d513fd8129d6894126702f578a33058
}
