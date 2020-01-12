package controllers

import (
	"WEB/models"
	//"fmt"
	"github.com/astaxie/beego"
)

type ModelsController struct {
	beego.Controller
}

func (c *ModelsController) Get() {
	user := models.UserInfo{Username: "liusi", Password: "12345678"}
	models.AddUser(&user)
	c.Ctx.WriteString("call model success!")
}
