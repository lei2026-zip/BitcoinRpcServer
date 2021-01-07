package db_mysql
//小鱼子

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
func ConnectDB(){
	config :=beego.AppConfig
	dbdriver :=config.String("db_driverName")
	//dbuser :=config.String("db_user")
	//dbpassword:=config.String("db_password")
	//dbip:=config.String("db_ip")
	//dbname :=config.String("db_name")
	connurl:= "db_user"+":"+"db_password"+"@tcp("+"db_ip"+")/"+"db_name"+"?charset=utf8"
	db,err:=sql.Open(dbdriver,connurl)
	if err !=nil{
		fmt.Println(err.Error())
		panic("数据库错误")
	}
	Db=db
}
