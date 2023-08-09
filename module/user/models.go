package user

import (
	"notebook/module/permission"
	"notebook/module/notes"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Fullname    string
	Password    string
	Avatar      string
	Email       string
	Notes				[]notes.Note
	Permissions []permission.Permission `gorm:"many2many:permission_users;"`
}
