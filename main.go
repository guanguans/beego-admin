package main

import (
	_ "beego-admin/routers"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "beego-admin/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
