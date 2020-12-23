package controllers

import (
	"BitcoinWeb/models/RPC"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strconv"
	"strings"
)

type RpcController struct {
	beego.Controller
}

type Commit struct {
	Url       string
	Commit    string
	Params    string
}

func (c *RpcController) Post() {
	commit := c.Ctx.Request.Body
	commitbyte, err := ioutil.ReadAll(commit)
	js := make(map[string]interface{})  //响应体map
	if err != nil {
		js["result"] = nil
		js["error"] = 1
		js["ErrorCode"] = err.Error()
		fmt.Println("数据解析错误!" + err.Error())
		str,err := json.Marshal(js)
		if err!=nil{
			fmt.Println("响应体序列化失败!" + err.Error())
			return
		}
		c.Ctx.ResponseWriter.Write(str)
		return
	}
	jsonmap := Commit{}
	json.Unmarshal(commitbyte, &jsonmap)
	fmt.Println(jsonmap.Commit)
	fmt.Println(jsonmap.Params)
	/**
	服务器发起本地rpc请求 获取并返回数据
	*/
	requst, err := RPC.NewRpcRequst()
	if err != nil {
		js["result"] = nil
		js["error"] = 2
		js["ErrorCode"] = err.Error()
		fmt.Println("rpc 新建错误!" + err.Error())
		str,err := json.Marshal(js)
		if err!=nil{
			fmt.Println("响应体序列化失败!" + err.Error())
			c.Ctx.ResponseWriter.Write(nil)
			return
		}
		c.Ctx.ResponseWriter.Write(str)
		return
	}
	//服务器发起rpc请求 执行客户端发来的命令

	result, err := requst.Rpc_DoPost(jsonmap.Commit,ParseParams(jsonmap.Params))
	if err != nil {
		js["result"] = nil
		js["error"] = 3
		js["ErrorCode"] =err.Error()
		fmt.Println("rpc 请求错误!" + err.Error())
		str,err := json.Marshal(js)
		if err!=nil{
			fmt.Println("响应体序列化失败!" + err.Error())
			c.Ctx.ResponseWriter.Write(nil)
			return
		}
		c.Ctx.ResponseWriter.Write(str)
		return
	}
	//序列化响应数据
	js["result"] = result
	js["error"] = 0
	js["ErrorCode"] = nil
	str,err := json.Marshal(js)
    if err!=nil{
		fmt.Println("响应体序列化失败!" + err.Error())
		c.Ctx.ResponseWriter.Write(nil)
		return
	}
	c.Ctx.ResponseWriter.Write(str)
}

/*
* 解析参数
* 转换为rpc的参数格式 [params...]
*/
func ParseParams(params string)[]interface{}{
	strarr :=  strings.Split(params," ")
	var inter []interface{}
	fmt.Println("len:",len(strarr))
	if len(params)>0 {
		for i := 0; i < len(strarr); i++ {
			in, err := strconv.Atoi(strarr[i])
			if err != nil {
				inter = append(inter, strarr[i])
				continue
			}
			inter = append(inter, in)
		}
	}
	fmt.Println("arr:",inter)
	return inter
}