package user

import (
	"BitcoinWeb/BTCUtils"
	"BitcoinWeb/BTCUtils/db_mysql"
	"encoding/hex"
)

type User struct {
	/*Id int `form:"id"`*/
	Name string `form:"phone"`
	Password string `form:"password"`
}

//使用sha256加密算法进行数据脱敏

func (u User) SaveUser()(int64,error){
	/*md5Hash :=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)*/
	u.Password = hex.EncodeToString(BTCUtils.Sha256Hassh([]byte(u.Password)))
	//数据库添加语句
	row,err := db_mysql.Db.Exec("insert into users(phone,password)values (?,?)",u.Name,u.Password)
	if err !=nil {
		return -1,err
	}
	id,err :=row.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}
/*func (u User)InsertUser()(int64, error){
	u.Password = hex.EncodeToString(BTCUtils.Sha256Hassh([]byte(u.Name)))
	row,err:= db_mysql.Db.Exec("insert into users(phone,password)values (?,?)",u.Name,u.Password)
	if err != nil {
		return -1,err
	}
	id,err:=row.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}*/
//查询已有的用户数据
func (u User) QueryUser()(error) {
	/*md5Hash :=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)*/
	u.Password = hex.EncodeToString(BTCUtils.Sha256Hassh([]byte(u.Password)))
	row:= db_mysql.Db.QueryRow("select phone from users where phone = ? and password = ?", u.Name,u.Password)
	err := row.Scan(&u.Name)
	if err != nil {
		return err
	}
	return nil
}
