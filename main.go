package main

import (
	_ "lxtkj/hellobeego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file",`{"filename":"logs/test.log"}`)
	beego.Run()
}

