package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	config.Configure()

	user := config.C.Database.User
	pass := config.C.Database.Password
	endpoint := config.C.Database.Endpoint
	dbname := config.C.Database.Name
	db, err = gorm.Open("mysql", user+":"+pass+"@("+endpoint+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.Pet{})
	db.AutoMigrate(&entity.Store{})
}
