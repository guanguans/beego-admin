package models

import (
	"time"
)

type AdminRole struct {
	Id        int       `orm:"column(id)" json:"id"`
	Name      string    `orm:"size(50);column(name)" json:"name"`
	Desc      string    `orm:"size(50);column(desc)" json:"desc"`
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp);column(created_at)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp);column(updated_at)" json:"updated_at"`
}
