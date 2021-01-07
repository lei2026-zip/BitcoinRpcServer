package controllers

import (
	"BitcoinWeb/models/RPC"
	"encoding/json"
	"errors"
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

/*
 2020年12月23日16:29:59
 处理 页面js 发起的post请求 并执行rpc请求 特定序列化返回结果
*/
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
	//pares params
	params,err := ParseParams(jsonmap.Params)
	if err!=nil{
		fmt.Println("Error:" + err.Error())
		c.Ctx.ResponseWriter.Write(nil)
		return
	}
	//服务器发起rpc请求 执行客户端发来的命令
	result, err := requst.Rpc_DoPost(jsonmap.Commit,params)
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
    st,err:=	json.Marshal(result)
	if err!=nil {
		fmt.Println(err)
		c.Ctx.ResponseWriter.Write(nil)
	}
	//序列化响应数据                                                       &nbsp; 转义 空格  &quot; 转义 "
 	js["result"] = Insert_Str(Insert_Str(Insert_Str(Insert_Str(Insert_Str(Insert_Str(string(st),",",",<br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"),`"`,"&quot;"),"{","{<br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"),"}","<br/>}"),`\n`,`<br/>`),`\`,``)
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
func ParseParams(params string)([]interface{},error){
	strarr :=  strings.Split(params," ")
	var inter []interface{}
	fmt.Println("len:",len(strarr))
	if len(params)>0 {
		for i := 0; i < len(strarr); i++ {
			if strings.Contains(strarr[i],`"`){
				str:= strings.Split(strarr[i],`"`)
				if(len(str)>=3){
					inter = append(inter,str[1])
				}else {
					return nil,errors.New("参数字符格式错误")
				}
			}else {
				in, err := strconv.Atoi(strarr[i])
				if err != nil {
					return nil,errors.New("参数整数格式错误")
				}
				inter = append(inter, in)
			}
		}
	}
	fmt.Println("arr:",inter)
	return inter,nil
}
/*
 插入字符串
*/
func Insert_Str(str string,sub string,new string) string{
	st := strings.Replace(str,sub,new,-1)
	return st
}
