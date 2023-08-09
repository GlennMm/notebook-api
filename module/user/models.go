package user

import (
	"notebook/module/notes"
	"notebook/module/permission"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Password string
	Avatar   string
	Email    string
	Notes    []notes.Note `gorm:"foreignkey:UserID"`
	Permissions []permission.Permission `gorm:"many2many:permission_users;"`
}
