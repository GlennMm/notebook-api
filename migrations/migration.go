package module

import (
	"notebook/module/notes"
	"notebook/module/permission"
	"notebook/module/user"

	"fmt"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&notes.Note{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&permission.Permission{})
	// if err != nil {
	// 	panic("Model miration failed. " + err.Error())
	// }
	fmt.Println("Model migrated successfully.")
}
