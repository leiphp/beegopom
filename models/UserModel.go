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

func UserStruct() []*UserModel {
	query := orm.NewOrm().QueryTable("user")
	data := make([]*UserModel,0)
	query.OrderBy("-user_id").All(&data)
	return data

}

func UserList(pageSize, page int) ([]*UserModel, int64){
	query := orm.NewOrm().QueryTable("user")
	total,_ := query.Count()

	offset := (page-1)*pageSize
	data := make([]*UserModel,0)
	query.OrderBy("-user_id").Limit(pageSize, offset).All(&data)

	return data,total
}


func GetUserByName(userkey string) UserModel{
	user := UserModel{UserKey:userkey}
	o := orm.NewOrm()
	o.Read(&user)
	return user
}