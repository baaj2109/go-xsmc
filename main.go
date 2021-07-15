package main

import (
	_ "hellobeego/routers"
	_ "hellobeego/sysinit"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}

/*

net start mysql80

# 新增資料庫
CREATE DATABASE `xcms`;


run server

bee run


*/
