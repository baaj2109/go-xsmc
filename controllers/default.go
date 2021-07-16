package controllers

import (
	"fmt"
	"hellobeego/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"

	//session使用
	c.SetSession("username", "baaj2109")
	user := c.GetSession("username")
	logs.Informational("user you loged in")
	fmt.Println(user)

	//更新邮箱
	models.UpdatePage()

	//查询数据并渲染
	m := models.GetPage()
	c.Data["Website"] = m.Website
	c.Data["Email"] = m.Email
	c.TplName = "index.tpl"
}
