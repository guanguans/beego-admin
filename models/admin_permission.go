package models

import (
	"time"
)

type AdminPermission struct {
	Id         int       `orm:"column(id)"`
	Name       string    `orm:"size(50);column(name)"`
	Slug       string    `orm:"size(50);column(slug)"`
	HttpMethod string    `orm:"size(255);column(http_method)"`
	HttpPath   string    `orm:"size(255);column(http_path)"`
	CreatedAt  time.Time `orm:"auto_now_add;type(timestamp);column(created_at)"`
	UpdatedAt  time.Time `orm:"auto_now;type(timestamp);column(updated_at)"`
}
