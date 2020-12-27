package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

var Db *sql.DB
func ConnectDB(){
	config :=beego.AppConfig
	dbdriver :=config.String("db_driverName")
	dbuser :=config.String("db_user")
	dbpassword:=config.String("db_password")
	dbip:=config.String("db_ip")
	dbname :=config.String("db_name")
	connurl:= dbuser+":"+dbpassword+"@tcp("+dbip+")/"+dbname+"?charset=utf8"
	db,err:=sql.Open(dbdriver,connurl)
	if err !=nil{
		fmt.Println(err.Error())
		panic("数据库错误")
	}
	Db=db
}
