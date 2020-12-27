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
		r.TplName = "error.html"
		return
	}
	err = user.QueryUser()
	if err !=nil  {
		r.TplName = "error.html"
		return
	}
	r.TplName ="Ahead.html"
}