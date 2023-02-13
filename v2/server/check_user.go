package main

import (
	md "draina/demo/v2/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func checkUser(id uint64) bool {
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to retrieve database")
	}
	db.AutoMigrate(md.User{})
	var user md.User
	res := db.First(&user, "uid = ?", id)
	return res.RowsAffected != 0
}
