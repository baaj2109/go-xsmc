package controllers

import (
	"fmt"
	"hellobeego/models"
	"strings"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	method := c.Ctx.Request.Method
	fmt.Println(method)
	if c.Ctx.Request.Method == "POST" {
		userkey := strings.TrimSpace((c.GetString("userkey")))
		password := strings.TrimSpace((c.GetString("password")))
		fmt.Println("userkey:" + userkey + ", password:" + password)
		if len(userkey) > 0 && len(password) > 0 {
			// password := utils.Md5([]byte(password))
			user := models.GetUserByName(userkey)
			// fmt.Println("user password " + user.PassWord + " password:" + password)
			if password == user.PassWord {
				c.SetSession("user", user)
				c.Redirect("/menu", 302)
				c.StopRun()
			}
		}
	}
	c.TplName = "login/index.html"
}
