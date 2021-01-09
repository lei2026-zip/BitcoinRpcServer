package controllers

import (
	"BitcoinRpcServer/models/RPC"
	"github.com/astaxie/beego"
)

type BitcoinToolController struct {
	beego.Controller
}

func (c *BitcoinToolController) Get(){
	c.Data["List"] = RPC.RPC_COMMIT
	c.TplName = "Ahead.html"
}

