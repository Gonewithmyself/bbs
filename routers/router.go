package routers

import (
	"bbs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/video", &controllers.MainController{}, "*:Play")

	beego.Router("/sfile", &controllers.MainController{}, "*:GetFiles")

	beego.Router("/table", &controllers.TableController{})
	beego.Router("/book", &controllers.BookController{})
}
