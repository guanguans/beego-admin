package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminUsers struct {
	Id            int       `orm:"column(id)"`
	Username      string    `orm:"size(190);unique;column(username)"`
	Password      string    `orm:"size(60);column(password)"`
	Name          string    `orm:"size(255);column(name)"`
	Avatar        string    `orm:"size(255);default();column(avatar)"`
	RememberToken string    `orm:"size(100);default();column(remember_token)"`
	CreatedAt     time.Time `orm:"column(created_at)"`
	UpdatedAt     time.Time `orm:"column(updated_at)"`
}

func (this *AdminUsers) GetAll() (adminUsers []AdminUsers) {
	orm.NewOrm().QueryTable(this).All(&adminUsers)

	return
}

func (this *AdminUsers) GetOne(id int) (adminUsers AdminUsers) {
	orm.NewOrm().QueryTable(this).Filter("id", id).One(&adminUsers)

	return
}
