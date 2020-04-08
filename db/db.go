package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/nfv-aws/wcafe-api-controller/config"
	"github.com/nfv-aws/wcafe-api-controller/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	conf := config.ReadDbConfig()
	db, err = gorm.Open("mysql", conf.User+":"+conf.Pass+"@("+conf.Endpoint+")/"+conf.Dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
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

	autoMigration()
}

func autoMigration() {
	db.AutoMigrate(&entity.Pet{})
	db.AutoMigrate(&entity.Store{})
}
