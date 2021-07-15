package routers

import (
	"hellobeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{}, "Get:Index")
	//menu
	beego.Router("/menu", &controllers.MenuController{}, "Get:Index")
	beego.Router("/menu/list", &controllers.MenuController{}, "*:List")
	beego.Router("/menu/edit", &controllers.MenuController{}, "*:Edit")
	beego.Router("/menu/editdo", &controllers.MenuController{}, "*:EditDo")
	beego.Router("/menu/add", &controllers.MenuController{}, "Get:Add")
	beego.Router("/menu/adddo", &controllers.MenuController{}, "*:AddDo")

	//user
	beego.Router("/user", &controllers.UserController{}, "Get:Index")
	beego.Router("/user/list", &controllers.UserController{}, "*:List")
	beego.Router("/user/add", &controllers.UserController{}, "Get:Add")
	beego.Router("/user/adddo", &controllers.UserController{}, "*:AddDo")
	beego.Router("/user/edit", &controllers.UserController{}, "Get:Edit")
	beego.Router("/user/editdo", &controllers.UserController{}, "*:EditDo")
	beego.Router("/user/deletedo", &controllers.UserController{}, "Get:DeleteDo")

	//login
	beego.Router("/login", &controllers.LoginController{}, "*:Index")

	//format
	beego.Router("/format/edit", &controllers.FormatController{}, "Get:Edit")
	beego.Router("/format/editdo", &controllers.FormatController{}, "*:EditDo")
}
