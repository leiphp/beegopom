package routers

import (
	"lxtkj/hellobeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/menu", &controllers.MenuController{}, "Get:Index")
	beego.Router("/menu/list", &controllers.MenuController{}, "*:List")
	beego.Router("/menu/edit", &controllers.MenuController{}, "*:Edit")
	beego.Router("/menu/editdo", &controllers.MenuController{}, "*:EditDo")
}
