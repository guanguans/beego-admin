package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlconn := beego.AppConfig.String("sqlconn")

	orm.RegisterDataBase("default", "mysql", sqlconn)

	orm.RegisterModel(
		new(AdminMenu),
		new(AdminOperationLog),
		new(AdminPermission),
		new(AdminRoleMenu),
		new(AdminRole),
		new(AdminRolePermission),
		new(AdminRoleUser),
		new(AdminUser),
		new(AdminUserPermission),
	)
}
