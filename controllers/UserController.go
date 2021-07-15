package controllers

import (
	"encoding/json"
	"hellobeego/consts"
	"hellobeego/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type UserController struct {
	BaseController
}

func (c *UserController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs.html"
	c.setTpl("user/index.html")
}

func (c *UserController) List() {
	result, count := models.UserList(10, 1)
	type UserEx struct {
		models.UserModel
		ParentName string
	}

	c.listJsonResult(consts.JRCodeSucc, "ok", count, result)
}

func (c *UserController) Add() {
	menu := models.ParentMenuList()
	menus := make(map[int]string)
	for _, v := range menu {
		menus[v.Mid] = v.Name
	}

	c.Data["Menus"] = menus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("user/add.html", "common/layout_edit.html")
}

func (c *UserController) AddDo() {
	password := strings.TrimSpace(c.GetString("Password"))
	password1 := strings.TrimSpace(c.GetString("Password1"))
	menu := models.ParentMenuList()
	//auth_str := []int{}这种方式初始化也行
	var auth_str []int
	for _, v := range menu {
		kint := v.Mid
		kstring := strconv.Itoa(kint) //int类型转string类型
		str := strings.TrimSpace(c.GetString("userauth_" + kstring))
		if str == "on" {
			auth_str = append(auth_str, v.Mid)
		}
	}

	var m models.UserModel
	if password == password1 {
		m.PassWord = password
	} else {
		return
	}
	//{{切片转成字符串
	strr := "["
	for k, v := range auth_str {
		if k < len(auth_str)-1 {
			strr = strr + strconv.Itoa(v) + ","
		} else {
			strr = strr + strconv.Itoa(v)
		}

	}
	strr = strr + "]"
	m.AuthStr = strr
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Insert(&m)
	}
}

func (c *UserController) Edit() {
	userId, _ := c.GetInt("userid")
	o := orm.NewOrm()
	var user = models.UserModel{UserId: userId}
	o.Read(&user)
	user.PassWord = ""
	c.Data["User"] = user

	authmap := make(map[int]bool)
	if len(user.AuthStr) > 0 {
		var authobj []int
		// json.Unmarshal([]byte(user.AuthStr), &authobj)
		str := []byte(user.AuthStr)
		json.Unmarshal(str, &authobj)
		for _, v := range authobj {
			authmap[v] = true
		}
	}
	type Menuitem struct {
		Name    string
		Ischeck bool
	}

	menu := models.ParentMenuList()
	menus := make(map[int]Menuitem)
	for _, v := range menu {
		menus[v.Mid] = Menuitem{v.Name, authmap[v.Mid]}
	}
	c.Data["Menus"] = menus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs_edit.html"
	c.setTpl("user/edit.html", "common/layout_edit.html")
}
func (c *UserController) EditDo() {
	password := strings.TrimSpace(c.GetString("Password"))
	password1 := strings.TrimSpace(c.GetString("Password1"))
	menu := models.ParentMenuList()
	//auth_str := []int{}这种方式初始化也行
	var auth_str []int
	for _, v := range menu {
		kint := v.Mid
		kstring := strconv.Itoa(kint) //int类型转string类型
		str := strings.TrimSpace(c.GetString("userauth_" + kstring))
		if str == "on" {
			auth_str = append(auth_str, v.Mid)
		}
	}
	var m models.UserModel
	if password == password1 {
		m.PassWord = password
	} else {
		return
	}
	//{{切片转成字符串
	strr := "["
	for k, v := range auth_str {
		if k < len(auth_str)-1 {
			strr = strr + strconv.Itoa(v) + ","
		} else {
			strr = strr + strconv.Itoa(v)
		}
	}
	strr = strr + "]"
	m.AuthStr = strr
	//}}
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Update(&m)
	}
}
func (c *UserController) DeleteDo() {

}
