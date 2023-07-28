package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/muxrestapi"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Photo{})

	DB = database
}
