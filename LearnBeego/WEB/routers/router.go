package routers

import (
	"WEB/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/testinput", &controllers.TestInputController{}, "get:Get;post:Post")
	beego.Router("/testlogin", &controllers.TestLoginController{}, "get:Login;post:Post")
	beego.Router("/testmodel", &controllers.TestModelController{}, "get:Get;post:Post")
	beego.Router("/models", &controllers.ModelsController{}, "get:Get;post:Post")
	beego.Router("/testview", &controllers.TestViewController{}, "get:Get;post:Post")
	beego.Router("/testhttplib", &controllers.TestHttpLibController{}, "get:Get;post:Post")
	beego.Router("/testcontext", &controllers.TestContextController{}, "get:Get;post:Post")
}
