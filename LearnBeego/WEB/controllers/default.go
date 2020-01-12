package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["name"] = "yanglt"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Ctx.WriteString("hello world!")
}
