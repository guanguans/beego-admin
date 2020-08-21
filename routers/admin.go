package routers

import (
	"beego-admin/controllers/AdminControllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("admin",
		beego.NSRouter("adminUsers", &AdminControllers.AdminUserController{}, "*:GetAll"),
		beego.NSRouter("adminUser/:id", &AdminControllers.AdminUserController{}, "*:GetOne"),

		beego.NSNamespace("auth",
			beego.NSRouter("login", &AdminControllers.AuthController{}, "post:Login"),
			beego.NSRouter("getToken", &AdminControllers.AuthController{}, "post:GetToken"),
			beego.NSRouter("verifyToken", &AdminControllers.AuthController{}, "post:VerifyToken"),
		),
	)
	beego.AddNamespace(ns)
}
