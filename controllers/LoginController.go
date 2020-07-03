package controllers

import (
	"github.com/astaxie/beego"
	"lxtkj/hellobeego/models"
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
			//password := utils.Md5([]byte(password))//通过utils工具md5密码加密
			user := models.GetUserByName(userkey)
			if password == user.PassWord{//跟数据库总加密后的密码做比较
				c.SetSession("user",user)
				c.Redirect("/",302)
				c.StopRun()
			}
		}
	}
	c.TplName = "login/index.html"
}