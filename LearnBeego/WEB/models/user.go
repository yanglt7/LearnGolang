package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)

type UserInfo struct {
	Id       int64
	Username string
	Password string
}

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:*hope8848@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))
	db = orm.NewOrm()
}

func AddUser(user_info *UserInfo) (int64, error) {
	id, err := db.Insert(user_info)
	return id, err
}

func ReadUserInfo(users *[]UserInfo) {
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("*").From("user_info")

	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}
