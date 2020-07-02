package controllers

import (
	"github.com/astaxie/beego"
	"lxtkj/hellobeego/models"
	"lxtkj/hellobeego/utils"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index(){
	if c.Ctx.Request.Method == "POST" {
		userkey := strings.TrimSpace(c.GetString("username"))
		password := strings.TrimSpace(c.GetString("password"))

		if len(userkey) > 0&& len(password) >0 {
			password := utils.Md5([]byte(password))
			user := models.GetUserByName(userkey)
			//if password == user.PassWord{
			if password == "123456" {
				c.SetSession("user",user)
				c.Redirect("/",302)
				c.StopRun()
			}
		}
	}
	c.TplName = "login/index.html"
}