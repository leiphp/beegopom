package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
	"lxtkj/hellobeego/consts"
	"lxtkj/hellobeego/models"
	"strconv"
	"time"
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

func (c *DataController) Edit() {
	did,_ := c.GetInt("did")
	if did>0{
		c.Data["Did"] = did
	}

	c.initForm(did)

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "data/footerjs_edit.html"
	c.setTpl("data/edit.html","common/layout_jfedit.html")
}

func (c *DataController) EditDo() {
	did,_ := c.GetInt("did")
	if did>0{
		if len(c.Ctx.Input.RequestBody) >0 {
			sj,err := simplejson.NewJson(c.Ctx.Input.RequestBody)
			if nil == err{
				var m models.DataModel
				m.Content = string(c.Ctx.Input.RequestBody)
				m.Did=did
				m.Parent=sj.Get("parent").MustInt()
				m.Mid=c.Mid
				m.Name=sj.Get("name").MustString()
				m.Seq=sj.Get("seq").MustInt()
				int, err := strconv.Atoi(sj.Get("status").MustString())//字符串转整数
				m.Status=int8(int)//int转int8
				m.UpdateTime = time.Now().Unix()
				id,err := orm.NewOrm().Update(&m)
				if nil ==err {
					c.jsonResult(consts.JRCodeSucc, "ok", id)
				}
			}
		}
	}
	c.jsonResult(consts.JRCodeFailed, "",0)
}

func (c *DataController) initForm(did int){
	format := models.MenuFormatStruct(c.Mid)
	if nil == format{
		return
	}
	schemaMap := format.Get("schema")
	formArray := format.Get("form")

	//初始化数据
	one := models.DataRead(did)
	if nil!=one{
		for k,_ := range schemaMap.MustMap(){
			switch schemaMap.GetPath(k, "type").MustString() {
			case "string":
				schemaMap.SetPath([]string{k,"default"}, one.Get(k).MustString())
				break
			case "integer":
				schemaMap.SetPath([]string{k,"default"}, one.Get(k).MustInt())
				break
			case "boolean":
				schemaMap.SetPath([]string{k,"default"}, one.Get(k).MustBool())
				break
			}
		}
	}
	//通用信息
	schemaMap.SetPath([]string{"parent","type"},"integer")
	schemaMap.SetPath([]string{"parent","title"},"上级数据")
	if nil!=one{
		schemaMap.SetPath([]string{"parent","default"},one.Get("parent").MustInt())
	}

	schemaMap.SetPath([]string{"name","type"},"string")
	schemaMap.SetPath([]string{"name","title"},"名称")
	if nil!=one{
		schemaMap.SetPath([]string{"name","default"},one.Get("name").MustString())
	}

	schemaMap.SetPath([]string{"seq","type"},"string")
	schemaMap.SetPath([]string{"seq","title"},"排序")
	if nil!=one{
		schemaMap.SetPath([]string{"seq","default"},one.Get("seq").MustInt())
	}
	schemaMap.SetPath([]string{"status","type"},"string")
	schemaMap.SetPath([]string{"status","title"},"状态")
	schemaMap.SetPath([]string{"status","enum"},[]string{"0","1"})
	if nil!=one{
		schemaMap.SetPath([]string{"status","default"},one.Get("status").MustInt())
	}
	c.Data["Schema"] = schemaMap.MustMap()

	//初始化通过form
	formarrayObj := formArray.MustArray()//formArray object
	if len(formarrayObj) <= 0 {
		var tmpArray []map[string]string
		tmpArray = append(tmpArray, map[string]string{"key":"parent"})
		tmpArray = append(tmpArray, map[string]string{"key":"name"})
		tmpArray = append(tmpArray, map[string]string{"key":"seq"})
		tmpArray = append(tmpArray, map[string]string{"key":"status"})

		for k,_ := range schemaMap.MustMap(){
			tmpArray = append(tmpArray, map[string]string{"key":k})
		}
		tmpArray = append(tmpArray, map[string]string{"type":"submit","title":"提交"})
		c.Data["Form"] = tmpArray
	}else{
		var tmpArray []interface{}
		tmpArray = append(tmpArray, map[string]string{"key":"parent"})
		tmpArray = append(tmpArray, map[string]string{"key":"name"})
		tmpArray = append(tmpArray, map[string]string{"key":"seq"})
		tmpArray = append(tmpArray, map[string]string{"key":"status"})

		var haveSubmit bool = false
		for k,v := range formArray.MustArray(){
			tmpArray = append(tmpArray,v)
			tmp:=formArray.GetIndex(k).Get("type")
			if "submit" == tmp.MustString(){
				haveSubmit= true
			}
		}
		if false == haveSubmit{
			tmpArray = append(tmpArray, map[string]string{"type":"submit","title":"提交"})
		}
		c.Data["Form"] = tmpArray
	}
}