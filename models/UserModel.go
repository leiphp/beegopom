package models

import "github.com/astaxie/beego/orm"

type UserModel struct {
	UserId int `orm:"pk;auto"`
	UserKey string `orm:"size(64);unique"`
	UserName string `orm:"size(64)"`
	AuthStr string `orm:"size(512)"`
	PassWord string `orm:"size(128)"`
	IsAdmin int8 `orm:"default(0)"`
}

func (m *UserModel) TableName() string {
	return "user"
}

func UserList(pageSize, page int)([]*UserModel, int64){
	query := orm.NewOrm().QueryTable("user")
	data := make([]*MenuModel,0)
	query.OrderBy("parent","-seq").All(&data)
	return nil,8
}