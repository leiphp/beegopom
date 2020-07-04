package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"sort"
	"github.com/bitly/go-simplejson"
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

func (m *MenuModel) TbNameMenu() string {
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

func MenuTreeStruct(user UserModel) map[int]MenuTree {
	//query := orm.NewOrm().QueryTable(TbNameMenu())
	query := orm.NewOrm().QueryTable("menu")
	data := make([]*MenuModel,0)
	query.OrderBy("parent","-seq").Limit(1000).All(&data)

	var menu = make(map[int]MenuTree)
	//auth
	if len(user.AuthStr)>0{
		var authArr []int
		json.Unmarshal([]byte(user.AuthStr), &authArr)
		sort.Ints(authArr)

		for _,v := range data{
			if 0==v.Parent {
				idx := sort.SearchInts(authArr, v.Mid)
				found := (idx<len(authArr) && authArr[idx] == v.Mid)
				if found{
					var tree = new (MenuTree)
					tree.MenuModel = *v
					menu[v.Mid] = *tree
				}
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
	query.OrderBy("-seq").Limit(1000).All(&data)
	return data
}

func MenuFormatStruct(mid int) *simplejson.Json {
	menu := MenuModel{Mid:mid}
	err := orm.NewOrm().Read(&menu)
	if nil == err {
		jsonstruct, err2 := simplejson.NewJson([]byte(menu.Format))
		if nil == err2 {
			return jsonstruct
		}
	}
	return nil
}
