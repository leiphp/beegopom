package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
)

type DataModel struct {
	Did int `orm:"pk;auto"`
	Mid int `orm:"default(0)"`
	Parent int `orm:"default(0)"`
	Name string `orm:"size(60)"`
	Content string `orm:"size(2048);default({})"`
	Seq int `orm:"index"`
	Status int8
	UpdateTime int64
}

func (m *DataModel) TableName() string {
	return "data"
}

func DataList(mid, pageSize, page int)([]map[string]interface{}, int64){
	if mid <= 0{
		return nil,0
	}
	//处理分页信息
	offset := (page-1)*pageSize
	query := orm.NewOrm().QueryTable("data").Filter("mid",mid)
	total,_ := query.Count()
	data := make([]*DataModel,0)
	query.OrderBy("parent","-seq").Limit(pageSize,offset).All(&data)
	dataEx := make([]map[string]interface{},0)
	for _,v :=range data{
		sj,err := simplejson.NewJson([]byte(v.Content))
		if nil==err{
			sj.Set("did",v.Did)
			sj.Set("name",v.Name)
			sj.Set("mid",v.Mid)
			sj.Set("parent",v.Parent)
			sj.Set("seq",v.Seq)
			sj.Set("status",v.Status)
			sj.Set("updatetime",v.UpdateTime)
			dataEx = append(dataEx,sj.MustMap())
		}
	}
	return dataEx,total
}

func DataRead(did int) *simplejson.Json{
	if did <= 0 {
		return nil
	}
	data := DataModel{Did:did}
	err := orm.NewOrm().Read(&data)
	if nil==err {
		sj, err2 := simplejson.NewJson([]byte(data.Content))
		if nil == err2 {
			sj.Set("did",data.Did)
			sj.Set("name",data.Name)
			sj.Set("mid",data.Mid)
			sj.Set("parent",data.Parent)
			sj.Set("seq",data.Seq)
			sj.Set("status",data.Status)
			sj.Set("updatetime",data.UpdateTime)

			return sj
		}
	}
	return nil
}