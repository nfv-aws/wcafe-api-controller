package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type Pet struct {

	// 商品の識別子
	Id string `gorm:"type:varchar(255)" json:"id"`

	// 商品種
	Species string `gorm:"NOT NULL;type:varchar(255)" json:"species" binding:"required"`

	// 商品名
	Name string `gorm:"NOT NULL;type:varchar(255)" json:"name" binding:"required"`

	// 商品の年齢
	Age int32 `gorm:"NOT NULL;type:int(11)" json:"age" binding:"gte=0,lt=200"`

	// 店の識別子
	StoreId string `gorm:"NOT NULL;type:varchar(255)" json:"store_id" binding:"required"`

	//作成日時
	CreatedAt time.Time `json:"created_time"`

	//更新日時
	UpdatedAt time.Time `json:"updated_time"`

	// ステータス
	Status string `json:"status"`
}

type PetRepository struct {
	DB *gorm.DB
}

func (pr *PetRepository) Find() ([]Pet, error) {
	log.Debug().Caller().Msg("pets Find")
	var r []Pet
	if err := pr.DB.Find(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func (pr *PetRepository) Create(p Pet) (Pet, error) {
	log.Debug().Caller().Msg("pets Create")
	if err := pr.DB.Create(p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (pr *PetRepository) Get(id string) (Pet, error) {
	log.Debug().Caller().Msg("pets Get")
	var r Pet
	if err := pr.DB.Where("id = ?", id).First(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func (pr *PetRepository) Update(id string, p Pet) (Pet, error) {
	log.Debug().Caller().Msg("pets Update")
	if err := pr.DB.Table("pets").Where("id = ?", id).Updates(p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (pr *PetRepository) Delete(id string) (Pet, error) {
	log.Debug().Caller().Msg("pets Delete")
	var r Pet
	if err := pr.DB.Table("pets").Where("id = ?", id).Delete(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}
