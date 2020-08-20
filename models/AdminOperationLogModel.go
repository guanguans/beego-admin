package models

import (
	"time"
)

type AdminOperationLog struct {
	Id        int       `orm:"column(id)"`
	UserId    int       `orm:"column(user_id)"`
	Path      string    `orm:"size(255);column(path)"`
	Method    string    `orm:"size(10);column(method)"`
	Ip        string    `orm:"size(255);column(ip)"`
	Input     string    `orm:"size(255);column(input)"`
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp);column(created_at)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp);column(updated_at)"`
}
