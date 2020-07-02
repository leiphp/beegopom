package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lxtkj/hellobeego/models"
	"lxtkj/hellobeego/consts"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName string
}

func (c *BaseController) Prepare()  {
	//赋值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
	//TODO 保存用户数据
	fmt.Println("beego:perpare"+c.controllerName+","+c.actionName)
	c.Data["Menu"] = models.MenuStruct()
}

//设置模板
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "common/layout.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(c.controllerName[0:len(c.controllerName) -10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName + "/" +actionName+".html"
	}
	_,found := c.Data["Footer"]
	if !found{
		c.Data["Footer"] = "menu/footerjs.html"
	}
	c.Layout = layout
	c.TplName = tplName
}

func (c *BaseController) jsonResult(code consts.JsonResultCode, msg string, obj interface{}){
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}