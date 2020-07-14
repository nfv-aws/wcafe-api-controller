package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type Store struct {

	// 店の識別子
	Id string `gorm:"type:varchar(255)" json:"id" `

	// 店名
	Name string `gorm:"type:varchar(255);NOT NULL;unique" json:"name" binding:"required"`

	// 店の種類
	Tag string `gorm:"type:varchar(255);NOT NULL" json:"tag" binding:"required"`

	// 住所
	Address string `gorm:"type:varchar(255);NOT NULL" json:"address" binding:"required"`

	// 店の強み
	StrongPoint string `gorm:"type:varchar(255)" json:"strong_point"`

	//作成日時
	CreatedAt time.Time `json:"created_time"`

	//更新日時
	UpdatedAt time.Time `json:"updated_time"`

	//ステータス
	Status string `json:"status"`
}

type StoreRepository struct {
	DB *gorm.DB
}

func (sr *StoreRepository) Find(limit int, offset int) ([]Store, error) {
	log.Debug().Caller().Msg("stores Find")
	log.Debug().Caller().Int("limit:", limit).Send()
	log.Debug().Caller().Int("offset:", offset).Send()
	var r []Store

	if err := sr.DB.Limit(limit).Offset(offset).Find(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func (sr *StoreRepository) Create(s Store) (Store, error) {
	log.Debug().Caller().Msg("stores Create")
	if err := sr.DB.Create(s).Error; err != nil {
		return s, err
	}
	return s, nil
}

func (sr *StoreRepository) Get(id string) (Store, error) {
	log.Debug().Caller().Msg("stores Get")
	var r Store

	if err := sr.DB.Where("id = ?", id).First(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func (sr *StoreRepository) Update(id string, s Store) (Store, error) {
	log.Debug().Caller().Msg("stores Update")
	if err := sr.DB.Table("stores").Where("id = ?", id).Updates(s).Error; err != nil {
		return s, err
	}
	return s, nil
}

func (sr *StoreRepository) Delete(id string) (Store, error) {
	log.Debug().Caller().Msg("stores Delete")
	var r Store
	if err := sr.DB.Table("stores").Where("id = ?", id).Delete(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func (sr *StoreRepository) PetsList(id string) ([]Pet, error) {
	log.Debug().Caller().Msg("stores PetsList")
	var s Store
	var p []Pet

	if err := sr.DB.Where("id = ?", id).First(&s).Error; err != nil {
		return p, err
	}
	if err := sr.DB.Table("pets").Where("store_id = ?", id).Find(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}
