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

	c.SetSession("cmusername", "Jack")

	user := c.GetSession(("cmusername"))
	logs.Informational("user Jack loged in ")
	fmt.Println(user)

	models.UpdatePage()
	m := models.GetPage()
	c.Data["Website"] = m.Website
	c.Data["Email"] = m.Email
	c.TplName = "index.tpl"
}
