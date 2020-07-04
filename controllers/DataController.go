package controllers

import (
	"lxtkj/hellobeego/consts"
	"lxtkj/hellobeego/models"
	"strconv"
)

type DataController struct {
	BaseController
	Mid int
}

func (c *DataController) Prepare() {
	//类似继承
	c.BaseController.Prepare()
	midstr := c.Ctx.Input.Param(":mid")
	c.Data["Mid"] = midstr
	mid,err := strconv.Atoi(midstr)
	if nil == err && mid > 0 {
		c.Mid = mid
	}else {
		c.setTpl()
	}
}

func (c *DataController) Index(){
	sj := models.MenuFormatStruct(c.Mid)
	if nil!=sj {
		title := make(map[string]string)
		titlemap := sj.Get("schema")
		for k,_ := range titlemap.MustMap(){
			stype := titlemap.GetPath(k,"type").MustString()
			if "object"!=stype && "array" !=stype{
				if len(titlemap.GetPath(k,"title").MustString())>0{
					title[k] = titlemap.GetPath(k,"title").MustString()
				}else{
					title[k] = k
				}
			}
		}
		c.Data["Title"] = title
	}
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "data/footerjs.html"
	c.setTpl()
}

func (c *DataController) List(){
	page,err := c.GetInt("page")
	if err != nil{
		page= 1
	}
	size,err:=c.GetInt("limit")
	if err != nil{
		size = 20
	}
	data,total := models.DataList(c.Mid,size,page)
	c.listJsonResult(consts.JRCodeSucc, "ok",total,data)
}