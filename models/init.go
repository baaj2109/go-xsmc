package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = true
	// orm.RegisterDataBase("default", "mysql", "root:tryit314641@tcp(127.0.0.1:3306)/xcms?charset=utf8")
	orm.RegisterModel(new(MenuModel))
	orm.RegisterModel(new(UserModel))
}
