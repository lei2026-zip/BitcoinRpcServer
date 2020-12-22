package controllers

import (
	"BitcoinWeb/models/user"
	"github.com/astaxie/beego"
)
type MainController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller

}

func (c *MainController) Get() {
	c.TplName = "register.html"
}
func (c *MainController) Post(){
	c.TplName = "register.html"
}

func (r *RegisterController) Get() {
	var user user.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("对不起，解析数据错误")
		return
	}
	_,err =user.SaveUser()
	if err !=nil  {
		r.Ctx.WriteString("对不起，用户注册失败")
		return
	}
	r.TplName ="login.html"
}

func (r *RegisterController) Post() {
	var user user.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("对不起，解析数据错误")
		return
	}
	_,err =user.SaveUser()
	if err !=nil  {
		r.Ctx.WriteString("对不起，用户注册失败")
		return
	}
	r.TplName ="login.html"
}
