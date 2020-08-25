package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminUser struct {
	Id        int       `orm:"column(id)" json:"id"`
	Username  string    `orm:"size(190);unique;column(username)" json:"username"`
	Password  string    `orm:"size(60);column(password)" json:"password"`
	Email     string    `orm:"size(60);column(email)" json:"email"`
	Phone     string    `orm:"size(60);column(phone)" json:"phone"`
	RealName  string    `orm:"size(60);column(real_name)" json:"real_name"`
	Nickname  string    `orm:"size(60);column(nickname)" json:"nickname"`
	Avatar    string    `orm:"size(255);default();column(avatar)" json:"avatar"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);column(created_at)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);column(updated_at)" json:"updated_at"`
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
