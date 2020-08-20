package models

import (
	"time"
)

type AdminUserPermission struct {
	Id           int       `orm:"column(id)"`
	UserId       int       `orm:"column(user_id)"`
	PermissionId int       `orm:"column(permission_id)"`
	CreatedAt    time.Time `orm:"auto_now_add;type(timestamp);column(created_at)"`
	UpdatedAt    time.Time `orm:"auto_now;type(timestamp);column(updated_at)"`
}
