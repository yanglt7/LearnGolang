package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type TestContextController struct {
	beego.Controller
}

func (c *TestContextController) Get() {
	c.Ctx.WriteString(c.Ctx.Input.IP() + ":" + strconv.Itoa(c.Ctx.Input.Port()))
	c.Ctx.WriteString(c.Ctx.Input.Query("name"))

	m := make(map[string]float64)
	m["zhangsan"] = 98.90
	c.Ctx.Output.JSON(m, false, false)
}

func (c *TestContextController) Post() {
	c.Ctx.WriteString("hello world!")
}
