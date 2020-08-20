package AdminControllers

import (
	"beego-admin/models"
)

type AdminUsersController struct {
	BaseController
}

var adminUsersModel models.AdminUsers

func (this *AdminUsersController) GetAll() {
	adminUsers := adminUsersModel.GetAll()
	this.Data["json"] = adminUsers
	this.ServeJSON()
}

func (this *AdminUsersController) GetOne() {
	id, _ := this.GetInt(":id")
	adminUsers := adminUsersModel.GetOne(id)
	this.Data["json"] = adminUsers
	this.ServeJSON()
}
