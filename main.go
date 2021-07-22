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

## open database
	net start mysql80

## install beego bee
	go get github.com/astaxie/beego
	go get github.com/astaxie/bee

## create web application
	bee new <project name>

## create api application
	bee apii <api name>

## run server
	bee run



*/
/*
獨立部屬

set GOOS = [ darwin(for mac) / linus/ windows]
set GOARCH = [ arm64 / amd64 ]
// go build >> xcms
bee pack >> xcms.tar.gz

nginx 雙機部屬 >> for load balance


*/
