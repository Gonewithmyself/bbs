package routers

import (
	"bbs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Home")
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.AutoRouter(&controllers.UserController{})
}
