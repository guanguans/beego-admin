package models

import (
	"time"
)

type AdminMenu struct {
	Id        int       `orm:"column(id)" json:"id"`
	ParentId  int       `orm:"column(parent_id)" json:"parent_id"`
	MenuName  string    `orm:"size(50);column(menu_name)" json:"menu_name"`
	MenuPath  string    `orm:"size(50);column(menu_path)" json:"menu_path"`
	Icon      string    `orm:"size(50);column(icon)" json:"icon"`
	IconSvg   string    `orm:"size(50);column(iconSvg)" json:"iconSvg"`
	IsShow    string    `orm:"size(50);column(is_show)" json:"is_show"`
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp);column(created_at)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp);column(updated_at)" json:"updated_at"`
}
