package models

import (
	"time"
)

type AdminRoleMenu struct {
	Id        int       `orm:"column(id)"`
	RoleId    int       `orm:"column(role_id)"`
	MenuId    int       `orm:"column(menu_id)"`
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp);column(created_at)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp);column(updated_at)"`
}
