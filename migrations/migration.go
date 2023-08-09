package module

import (
	"notebook/module/notes"
	"notebook/module/permission"
	"notebook/module/user"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{}, &notes.Note{}, &permission.Permission{})
	if err != nil {
		panic("Model miration failed. " + err.Error())
	}
}
