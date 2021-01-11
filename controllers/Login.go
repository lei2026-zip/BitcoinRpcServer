package controllers

import (
	"BitcoinRpcServer/models/RPC"
	"BitcoinRpcServer/models/user"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (r *LoginController) Post(){
	var user user.User
	err :=r.ParseForm(&user)
	if err != nil {
		fmt.Println("对不起，解析数据错误",err)
		r.TplName ="error.html"
		return
	}
	err = user.QueryUser()
	if err !=nil  {
		fmt.Println("对不起，数据查询失败",err)
		r.TplName ="error.html"
		return
	}
	r.Data["List"] = RPC.RPC_COMMIT
	r.TplName ="Ahead.html"
}