package models

type AdminPermission struct {
	Id     int `orm:"column(id)" json:"id"`
	MenuId int `orm:"column(menu_id)" json:"menu_id"`
	RoleId int `orm:"column(role_id)" json:"role_id"`
}
