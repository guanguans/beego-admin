package admin

import (
	"beego-admin/models"
)

type AdminUserController struct {
	AuthController
}

var adminUserModel models.AdminUser

func (c *AdminUserController) GetAll() {
	c.AjaxSuccess(adminUserModel.GetAll())
}

func (c *AdminUserController) GetOne() {
	id, _ := c.GetInt(":id")
	c.AjaxSuccess(adminUserModel.GetOne(id))
}
