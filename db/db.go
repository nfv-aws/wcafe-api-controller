package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"

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

	user := config.C.DB.User
	pass := config.C.DB.Password
	endpoint := config.C.DB.Endpoint
	dbname := config.C.DB.Name
	db, err = gorm.Open("mysql", user+":"+pass+"@("+endpoint+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error().Caller().Err(err).Send()
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
	db.AutoMigrate(&entity.Store{})
	db.AutoMigrate(&entity.Pet{}).AddForeignKey("store_id", "stores(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&entity.User{})
}
