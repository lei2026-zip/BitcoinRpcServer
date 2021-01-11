package user
//小鱼子
import (
	"BitcoinRpcServer/BTCUtils"
	"BitcoinRpcServer/BTCUtils/db_mysql"
	"encoding/hex"
)

type User struct {
	/*Id int `form:"id"`*/
	Name string `form:"name"`
	Password string `form:"password"`
}
type NewUser struct {
	NewName string `form:"newName"`
	NewPassword string `form:"newPassword"`
	NewE_mail string `form:"newE_mail"`
}
//使用sha256加密算法进行数据脱敏

func (u NewUser) SaveUser()(int64,error){
	/*md5Hash :=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)*/
	u.NewPassword = hex.EncodeToString(BTCUtils.Sha256Hassh([]byte(u.NewPassword)))
	//数据库添加语句
	row,err := db_mysql.Db.Exec("insert into users(name,password,email)values (?,?,?)",u.NewName,u.NewPassword,u.NewE_mail)
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
	row:= db_mysql.Db.QueryRow("select name from users where name = ? and password = ?", u.Name,u.Password)
	err := row.Scan(&u.Name)
	if err != nil {
		return err
	}
	return nil
}
