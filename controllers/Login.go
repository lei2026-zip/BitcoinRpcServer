package controllers

import (
	"BitcoinWeb/models/user"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (r *LoginController) Post(){
	var user user.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("对不起，解析数据错误")
		r.TplName ="error.html"
		return
	}
	err = user.QueryUser()
	if err !=nil  {
		r.Ctx.WriteString("对不起，数据查询失败")
		r.TplName ="error.html"
		return
	}
	r.TplName ="Ahead.html"
}