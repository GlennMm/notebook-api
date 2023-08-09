package notes

import (
	"notebook/module/permission"

	"gorm.io/gorm"
)

type Note struct {
	*gorm.Model
	UserId      uint
	Text        string
	Permissions []permission.Permission
	State       string
}
