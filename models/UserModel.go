package models

import "github.com/astaxie/beego/orm"

type UserModel struct {
	UserId   int    `orm:"pk;auto"`
	UserKey  string `orm:"size(64);unique"`
	UserName string `orm:"size(64)"`
	AuthStr  string `orm:"size(512)"`
	PassWord string `orm:"size(128)"`
	IsAdmin  int8   `orm:"default(0)"`
}

func (m *UserModel) TableName() string {
	return "xcms_user"
}

func UserStruct() []*UserModel {
	query := orm.NewOrm().QueryTable("xcms_user")
	data := make([]*UserModel, 0)
	query.OrderBy("-user_id").All(&data)
	return data

}

func UserList(pageSize, page int) ([]*UserModel, int64) {
	query := orm.NewOrm().QueryTable("xcms_user")
	data := make([]*MenuModel, 0)
	query.OrderBy("parent", "-seq").All(&data)
	return nil, 8
}

func GetUserByName(userkey string) UserModel {
	o := orm.NewOrm()
	user := UserModel{UserKey: userkey}
	o.Read(&user, "user_key")
	return user
}
