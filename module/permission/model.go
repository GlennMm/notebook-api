package permission

import (
	"os/user"

	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	NoteId int
	Users  []user.User `gorm:"many2many:permission_users;"`
}
