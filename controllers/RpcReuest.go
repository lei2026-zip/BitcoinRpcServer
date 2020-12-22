package controllers

import (
	"github.com/astaxie/beego"
)

type RpcController struct {
	beego.Controller
}


func (c *RpcController) Post(){
    err :=  c.Ctx.Request.ParseForm()
	if err!=nil {
		c.Ctx.WriteString("数据解析错误!")
	}

	/**
	    服务器发起本地rpc请求 获取并返回数据
	 */
	 //requst,err := RPC.NewRpcRequst()
	 if err!=nil{
		c.Ctx.WriteString("rpc 请求错误!")
	 }
	// requst.Rpc_DoPost()
}
