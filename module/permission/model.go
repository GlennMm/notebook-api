package permission

import (
	"os/user"

	"gorm.io/gorm"
)

type Permission struct {
	*gorm.Model
	NoteId   uint
	UsersIds []user.User `gorm:"many2many:permission_users;"`
}
