package models

import (
	"time"
)

type AdminRole struct {
	Id        int       `orm:"column(id)"`
	Name      string    `orm:"size(50);column(name)"`
	Slug      string    `orm:"size(50);column(slug)"`
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp);column(created_at)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp);column(updated_at)"`
}
