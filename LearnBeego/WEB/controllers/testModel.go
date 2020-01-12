package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TestModelController struct {
	beego.Controller
}

//UserInfo -> user_info
type UserInfo struct {
	Id       int64
	Username string
	Password string
}

func (c *TestModelController) Get() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:*hope8848@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()
	//下面是插入
	//user := UserInfo{Username: "zhangsan", Password: "123456"}
	//id, err := o.Insert(&user)

	//下面是更新
	// user := UserInfo{Username: "zhangsan", Password: "123456"}
	// user.Id = 1
	// user.Username = "lisi"
	// id, err := o.Update(&user)

	//下面是读取
	// user := UserInfo{Username: "zhangsan", Password: "123456"}
	// user.Id = 1
	// o.Read(&user)

	//下面是原生读取
	//var maps []orm.Params
	//o.Raw("select * from user_info").Values(&maps)

	// for _, v := range maps {
	// 	c.Ctx.WriteString(fmt.Sprintf("user_info: %v", v))
	// }

	// var users []UserInfo
	// o.Raw("select username from user_info").QueryRows(&users)
	// c.Ctx.WriteString(fmt.Sprintf("user_info: %v", users))

	//采用queryBuilder
	var users []UserInfo
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("password").From("user_info").Where("username='lisi'").And("id=1").Limit(1)

	sql := qb.String()
	o.Raw(sql).QueryRows(&users)

	c.Ctx.WriteString(fmt.Sprintf("user_info: %v", users))
}
