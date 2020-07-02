package models

import (
	"github.com/astaxie/beego/orm"
)

type MenuModel struct {
	Mid int `orm:"pk;auto"`
	Parent int
	Name string `orm:"size(45)"`
	Seq int
	Format string `orm:"size(2048);default({})"`
}
type MenuTree struct {
	MenuModel
	Child []MenuModel
}

func (m *MenuModel) TableName() string {
	return "menu"
}

func MenuStruct() map[int]MenuTree {
	query := orm.NewOrm().QueryTable("menu")
	data := make([]*MenuModel,0)
	query.OrderBy("parent","-seq").All(&data)

	var menu = make(map[int]MenuTree)
	if(len(data)>0){
		for _,v := range data{
			if 0==v.Parent {
				var tree = new (MenuTree)
				tree.MenuModel = *v
				menu[v.Mid] = *tree
			}else{
				if tmp,ok := menu[v.Parent];ok{
					tmp.Child = append(tmp.Child, *v)
					menu[v.Parent] = tmp
				}
			}
		}
	}
	return menu
}

func MenuList() ([] *MenuModel, int64){
	query := orm.NewOrm().QueryTable("menu")
	total,_ := query.Count()
	data := make([]*MenuModel, 0)
	query.OrderBy("parent", "-seq").All(&data)
	return data,total
}

func ParentMenuList() []*MenuModel {
	query := orm.NewOrm().QueryTable("menu").Filter("parent",0)
	data := make([]*MenuModel,0)
	query.OrderBy("-seq").All(&data)
	return data
}

