package routers

import (
	"BitcoinWeb/controllers"
	"github.com/astaxie/beego"
)


func init() {
	//发送登录界面
    beego.Router("/Bitcoin", &controllers.LoginController{})
    //获取主页
	beego.Router("/BitcoinTool", &controllers.BitcoinToolController{})
    //获取页面发起的rpc post请求
	beego.Router("/RPC", &controllers.RpcController{})
}
