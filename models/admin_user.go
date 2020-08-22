package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminUser struct {
	Id            int       `orm:"column(id)"`
	Username      string    `orm:"size(190);unique;column(username)"`
	Password      string    `orm:"size(60);column(password)"`
	Name          string    `orm:"size(255);column(name)"`
	Avatar        string    `orm:"size(255);default();column(avatar)"`
	RememberToken string    `orm:"size(100);default();column(remember_token)"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime);column(created_at)"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime);column(updated_at)"`
}

func (this *AdminUser) GetByUsername(userName string) (adminUser AdminUser) {
	orm.NewOrm().QueryTable(this).Filter("username", userName).One(&adminUser)
	return
}

func (this *AdminUser) GetAll() (adminUsers []AdminUser) {
	orm.NewOrm().QueryTable(this).All(&adminUsers)
	return
}

func (this *AdminUser) GetOne(id int) (adminUser AdminUser) {
	orm.NewOrm().QueryTable(this).Filter("id", id).One(&adminUser)
	return
}
