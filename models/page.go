package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Id      int
	Website string
	Email   string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8",30)
	orm.RegisterModel(new(Page))
}
/*获取数据*/
func GetPage() Page {
	//rtn := Page{Website:"leixiaotian.cn",Email:"1124378213@qq.com"}
	//return rtn
	o := orm.NewOrm()
	p := Page{Id: 1}
	err := o.Read(&p)
	fmt.Println(err)

	return p
}

/*更新数据*/
func UpdatePage(){
	p := Page{Id:1,Email:"myemail_updated"}
	o := orm.NewOrm()
	o.Update(&p,"Email")
}