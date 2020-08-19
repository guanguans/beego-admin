package AdminControllers

import (
	"beego-admin/models"
)

type AdminUsersController struct {
	BaseController
}

func (this*AdminUsersController)GetAll()  {
	var adminUsersModel models.AdminUsers
	adminUsers := adminUsersModel.GetAll()
	this.Data["json"] = adminUsers
	this.ServeJSON()
}

func (this*AdminUsersController)GetOne()  {
	id, _ := this.GetInt(":id")
	var adminUsersModel models.AdminUsers
	adminUsers := adminUsersModel.GetOne(id)
	this.Data["json"] = adminUsers
	this.ServeJSON()
}