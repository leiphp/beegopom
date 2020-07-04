package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"lxtkj/hellobeego/consts"
	"lxtkj/hellobeego/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs.html"
	c.setTpl("user/index.html")
}

func (c *UserController) List() {
	result,count := models.UserList(10, 1)
	type UserEx struct {
		models.UserModel
		ParentName string
	}

	c.listJsonResult(consts.JRCodeSucc, "ok", count, result)
}

func (c *UserController) Add(){
	menu := models.ParentMenuList()
	fmt.Println(menu)
	menus := make(map[int]string)
	for _,v := range menu{
		menus[v.Mid] = v.Name
	}

	c.Data["Menus"] = menus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("user/add.html","common/layout_edit.html")
}

func (c *UserController) AddDo(){
	var m models.UserModel
	if err := c.ParseForm(&m); err==nil{
		orm.NewOrm().Insert(&m)
	}
}

func (c *UserController) Edit(){
	userId,_ := c.GetInt("userid")
	o := orm.NewOrm()
	var user = models.UserModel{UserId:userId}
	o.Read(&user)
	user.PassWord = ""
	c.Data["User"] = user

	authmap := make(map[int]bool)
	if len(user.AuthStr) >0 {
		var authobj []int
		str := []byte(user.AuthStr)//user.AuthStr格式"[1,5]"
		json.Unmarshal(str, &authobj)
		fmt.Println(authobj)//[1.5]切片
		for _,v := range authobj {
			authmap[v] = true
		}
		fmt.Println(authmap)//map[1:true 5:true]
	}
	type Menuitem struct {
		Name string
		Ischeck bool
	}

	menu := models.ParentMenuList()
	menus := make(map[int]Menuitem)
	for _,v := range menu{
		menus[v.Mid] = Menuitem{v.Name,authmap[v.Mid]}
	}
	c.Data["Menus"] = menus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("user/edit.html","common/layout_edit.html")
}

func (c *UserController) EditDo(){

}

func (c *UserController) DeleteDo(){

}