package routers

import (
	"BitcoinRpcServer/controllers"
	"github.com/astaxie/beego"
)


func init() {
	//先执行“/”打开注册页面
	beego.Router("/", &controllers.MainController{})
	//执行“register”打开登录页面
	beego.Router("/register", &controllers.RegisterController{})
	//执行“login”打开主页面
	beego.Router("/login",&controllers.LoginingController{})
	beego.Router("/Bitcoin", &controllers.RpcController{})
}
