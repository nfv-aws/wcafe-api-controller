package sqlmock

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	// 	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
)

// 空のmockを作成
func NewMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	um, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return um, mock, nil
}
