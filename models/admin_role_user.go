package models

type AdminRoleUser struct {
	Id     int `orm:"column(id)" json:"id"`
	UserId int `orm:"column(user_id)" json:"user_id"`
	RoleId int `orm:"column(role_id)" json:"role_id"`
}
