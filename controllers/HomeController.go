package controllers

import (

)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"

	//查询数据并渲染
	//c.TplName = "user/index.html"
	c.setTpl("home/index.html")
}

