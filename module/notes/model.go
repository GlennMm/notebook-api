package notes

import (
	"notebook/module/permission"

	"gorm.io/gorm"
)

type Note struct {
	*gorm.Model
	UserID      int
	Text        string
	Permissions []permission.Permission
	State       string
}
