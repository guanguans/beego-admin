package AdminControllers

import (
	"beego-admin/models"
)

type AdminUserController struct {
	BaseController
}

var adminUserModel models.AdminUser

func (this *AdminUserController) GetAll() {
	adminUsers := adminUserModel.GetAll()
	this.Data["json"] = adminUsers
	this.ServeJSON()
}

func (this *AdminUserController) GetOne() {
	id, _ := this.GetInt(":id")
	adminUser := adminUserModel.GetOne(id)
	this.Data["json"] = adminUser
	this.ServeJSON()
}
