package controllers

import (
	"WEB/models"

	"github.com/astaxie/beego"
)

type TestViewController struct {
	beego.Controller
}

func (c *TestViewController) Get() {
	var users []models.UserInfo
	models.ReadUserInfo(&users)

	c.Data["Users"] = users
	c.Data["len"] = len(users)

	// c.Data["IsDisplay"] = false
	// c.Data["Title"] = "hello"
	// c.Data["Content"] = "yanglt"
	// c.Data["Content2"] = "Y_momo"
	c.TplName = "test_view.tpl"
}
