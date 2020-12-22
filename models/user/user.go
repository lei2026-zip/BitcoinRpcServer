package user

import (
	"BitcoinRpcServer/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	password string `form:"password"`
}
func (u User) SaveUser()(int64,error){
	md5Hash :=md5.New()
	md5Hash.Write([]byte(u.password))
	passwordBytes:=md5Hash.Sum(nil)
	u.password = hex.EncodeToString(passwordBytes)
	//数据库添加语句
	row,err :=db_mysql.Db.Exec("insert into users(phone,password)values (?,?)",u.Phone,u.password)
	if err !=nil {
		return -1,err
	}
	id,err :=row.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}
//查询已有的用户数据
func (u User) QueryUser()(error) {
	md5Hash :=md5.New()
	md5Hash.Write([]byte(u.password))
	passwordBytes := md5Hash.Sum(nil)
	u.password = hex.EncodeToString(passwordBytes)
	row:=db_mysql.Db.QueryRow("select phone from users where phone = ? and password = ?", u.Phone,u.password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return err
	}
	return nil
}
