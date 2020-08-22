package routers

import (
	"beego-admin/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("admin",
		beego.NSRouter("login", &admin.LoginController{}, "post:Login"),
		beego.NSRouter("getAdminUser", &admin.AuthController{}, "post:GetAdminUser"),

		beego.NSNamespace("adminUser",
			beego.NSRouter("index", &admin.AdminUserController{}, "*:GetAll"),
			beego.NSRouter("show/:id", &admin.AdminUserController{}, "*:GetOne"),
		),
	)
	beego.AddNamespace(ns)
}
