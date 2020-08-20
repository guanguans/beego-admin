package models

import (
	"time"
)

type AdminMenu struct {
	Id         int       `orm:"column(id)"`
	ParentId   int       `orm:"column(parent_id)"`
	Order      int       `orm:"column(order)"`
	Title      string    `orm:"size(50);column(title)"`
	Icon       string    `orm:"size(50);column(icon)"`
	Uri        string    `orm:"size(255);column(uri)"`
	Permission string    `orm:"size(255);column(permission)"`
	CreatedAt  time.Time `orm:"auto_now_add;type(timestamp);column(created_at)"`
	UpdatedAt  time.Time `orm:"auto_now;type(timestamp);column(updated_at)"`
}
