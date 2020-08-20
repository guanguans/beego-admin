package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlconn := beego.AppConfig.String("sqlconn")
	orm.RegisterDataBase("default", "mysql", sqlconn)
	orm.RegisterModel(new(AdminUser))
}
