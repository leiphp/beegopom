package models

import "github.com/astaxie/beego/orm"

func init() {
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8",30)//initDb.go已经配置了
	orm.RegisterModel(new(MenuModel))
	orm.RegisterModel(new(UserModel))
	orm.RegisterModel(new(DataModel))//第一次初始化表之后可以注销掉
}
