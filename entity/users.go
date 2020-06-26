package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type User struct {

	// ユーザの識別子
	Id string `gorm:"type:varchar(255)" json:"id"`

	// 会員番号
	Number int32 `gorm:"type:int(32) AUTO_INCREMENT;NOT NULL;unique" json:"number"  binding:"required"`

	// 氏名
	Name string `gorm:"type:varchar(255)" json:"name"`

	// 住所
	Address string `gorm:"type:varchar(255)" json:"address"`

	// メールアドレス
	Email string `gorm:"type:varchar(255)" json:"email" validate:"omitempty,email"`

	//作成日時
	CreatedAt time.Time `json:"created_time"`

	//更新日時
	UpdatedAt time.Time `json:"updated_time"`
}

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) Find() ([]User, error) {
	log.Debug().Caller().Msg("users Find")
	var r []User
	if err := ur.DB.Find(&r).Error; err != nil {
		return r, err
	}

	return r, nil
}

func (ur *UserRepository) Create(u User) (User, error) {
	log.Debug().Caller().Msg("users Create")
	if err := ur.DB.Create(u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (ur *UserRepository) Get(id string) (User, error) {
	log.Debug().Caller().Msg("users Get")
	var r User

	if err := ur.DB.Where("id = ?", id).First(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func (ur *UserRepository) Update(id string, u User) (User, error) {
	log.Debug().Caller().Msg("users Update")
	if err := ur.DB.Table("users").Where("id = ?", id).Updates(u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (ur *UserRepository) Delete(id string) (User, error) {
	log.Debug().Caller().Msg("users Delete")
	var r User
	if err := ur.DB.Table("users").Where("id = ?", id).Delete(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}
