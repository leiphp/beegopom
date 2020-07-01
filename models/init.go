package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8",30)
	orm.RegisterModel(new(MenuModel))
}
