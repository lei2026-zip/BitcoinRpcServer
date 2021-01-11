package controllers

import (
	"BitcoinRpcServer/models/BitcoinAddress"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	_"BitcoinRpcServer/models/BitcoinAddress"
)

type VerifyController struct {
	beego.Controller
}

func (c *VerifyController) Post(){
	commit := c.Ctx.Request.Body
	commitbyte, err := ioutil.ReadAll(commit)
	js := make(map[string]interface{})  //响应体map
	json_reust := make(map[string]string)
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
    json.Unmarshal(commitbyte,&json_reust)
	fmt.Println("requst_target:"+json_reust["requst_target"])
	if json_reust["requst_target"]=="verify"{
		flag := BitcoinAddress.CheckAddress(json_reust["params"])
		fmt.Println("Verify:",flag)
		if flag{
			js["result"] = "验证通过!."
		}else{
			js["result"] = "验证错误!"
		}
	}else if json_reust["requst_target"]== "NewBitcoin"{
		  pubk := BitcoinAddress.New_elliptic_publickey()
		  if pubk==nil{
		  	  js["result"] = "生成错误!"
			  js["error"] = 1
			  js["ErrorCode"] = nil
			  str,err := json.Marshal(js)
			  if err!=nil{
				  fmt.Println("响应体序列化失败!" + err.Error())
				  c.Ctx.ResponseWriter.Write(nil)
				  return
			  }
			  c.Ctx.ResponseWriter.Write(str)
		  }
		  addr := BitcoinAddress.NewBitcoinAddress(0x00,pubk)
		  js["result"] = addr
	}else {
		js["result"]=nil
	}
	fmt.Println(js["result"])
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

