package controllers

import (
	"hellobeego/consts"
	"hellobeego/models"

	"github.com/astaxie/beego/orm"
)

type MenuController struct {
	BaseController
}
type MenuEx struct {
	models.MenuModel
	ParentName string
}

func (c *MenuController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs.html"
	c.setTpl("menu/index.html")
	// c.setTp()
}

func (c *MenuController) List() {
	data, total := models.MenuList()

	menu := make(map[int]string)
	for _, v := range data {
		menu[v.Mid] = v.Name
	}

	var dataEx []MenuEx
	for _, v := range data {
		dataEx = append(dataEx, MenuEx{*v, menu[v.Parent]})
	}
	c.listJsonResult(consts.JRCodeSucc, "ok", total, data)
}

func (c *MenuController) Add() {
	// not yet
	var pMenus []models.MenuModel
	data, _ := models.MenuList()
	for _, v := range data {
		if 0 == v.Parent {
			pMenus = append(pMenus, *v)
		}
	}
	c.Data["PMenus"] = pMenus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs_edit.html"
	c.setTpl("menu/add.html", "common/layout_edit.html")
}

func (c *MenuController) AddDo() {
	//not yet
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Insert(&m)
	}
}

func (c *MenuController) Edit() {
	c.Data["Mid"] = c.GetString("mid")
	c.Data["Parent"], _ = c.GetInt("parent")
	c.Data["Seq"] = c.GetString("seq")
	c.Data["Name"] = c.GetString("name")

	var pMenus []models.MenuModel
	data, _ := models.MenuList()
	for _, v := range data {
		if 0 == v.Parent {
			pMenus = append(pMenus, *v)
		}
	}
	c.Data["PMenus"] = pMenus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs_edit.html"
	c.setTpl("menu/edit.html", "common/layout_edit.html")
}

func (c *MenuController) EditDo() {
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Update(&m)
	}
}
