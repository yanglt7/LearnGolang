package controllers

import (
	"github.com/astaxie/beego"
)

type TestLoginController struct {
	beego.Controller
}

type UserData struct {
	Username string
	Password string
}

func (c *TestLoginController) Login() {
	name := c.Ctx.GetCookie("name")
	password := c.Ctx.GetCookie("password")

	//c.Ctx.SetCookie("name", name, -1, "/")
	//c.Ctx.SetCookie("password", password, -1, "/")

	if name != "" {
		c.Ctx.WriteString("Username:" + name + "Password:" + password)
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/testlogin" method="post">
                        <input type="text" name="Username"/>
                        <input type="password" name="Password"/>
                        <input type="submit" value="提交"/>
                        </form></html>`)
	}

}

func (c *TestLoginController) Post() {
	u := UserData{}
	if err := c.ParseForm(&u); err != nil {

	}

	c.Ctx.SetCookie("name", u.Username, 100, "/")
	c.Ctx.SetCookie("password", u.Password, 100, "/")
	c.SetSession("name", u.Username)
	c.SetSession("password", u.Password)
	c.Ctx.WriteString("Username:" + u.Username + "Password:" + u.Password)
}
