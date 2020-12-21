package routers

import (
	"BitcoinRpcServer/controllers"
	"github.com/astaxie/beego"
)


func init() {
<<<<<<< HEAD
	//发送登录界面
    beego.Router("/Bitcoin", &controllers.LoginController{})
    //获取主页
	beego.Router("/BitcoinTool", &controllers.BitcoinToolController{})
    //获取页面发起的rpc post请求
	beego.Router("/RPC", &controllers.RpcController{})
=======
	//先执行“/”打开注册页面
	beego.Router("/", &controllers.MainController{})
	//执行“register”打开登录页面
	beego.Router("/register", &controllers.RegisterController{})
	//执行“login”打开主页面
	beego.Router("/login",&controllers.LoginingController{})
	beego.Router("/Bitcoin", &controllers.RpcController{})
>>>>>>> 440a064dc656d1e5482cfe564023c095303082ae
}
