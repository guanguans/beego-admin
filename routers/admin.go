package routers

import (
	"beego-admin/controllers/AdminControllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("admin",
		beego.NSRouter("adminUsers", &AdminControllers.AdminUsersController{}, "*:GetAll"),
		beego.NSRouter("adminUsers/:id", &AdminControllers.AdminUsersController{}, "*:GetOne"),
	)
	beego.AddNamespace(ns)
}
