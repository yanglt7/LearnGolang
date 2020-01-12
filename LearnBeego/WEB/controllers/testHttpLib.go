package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type TestHttpLibController struct {
	beego.Controller
}

func (c *TestHttpLibController) Get() {
	req := httplib.Get("http://douban.com")
	str, err := req.String()

	if err != nil {
		panic(err)
	}

	c.Ctx.WriteString(str)
}

func (c *TestHttpLibController) Post() {
	c.Ctx.WriteString("hello world!")
}
