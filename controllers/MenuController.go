package controllers

import (
	//"github.com/astaxie/beego/orm"
	//"lxtkj/hellobeego/models"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs.html"
	c.setTpl("menu/index.html")
}

func (c *MenuController) List() {

}

func (c *MenuController) Edit() {

}

func (c *MenuController) EditDo() {

}