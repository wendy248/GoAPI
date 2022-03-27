package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/dbtugas?charset=utf8&parseTime=True&loc=Local") //type of database, name:pass
	if err != nil {
		panic("connection to database error")
	}

	db.AutoMigrate(&Mahasiswa{}, &MataKuliah{})	//untuk menambah tabel
	return db
}
