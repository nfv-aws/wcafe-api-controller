package entity

import (
	"github.com/jinzhu/gorm"
	"time"
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

func (p *PetRepository) Find() ([]Pet, error) {
	var r []Pet
	if err := p.DB.Find(&r).Error; err != nil {
		return r, err
	}

	return r, nil
}

func (p *PetRepository) Create(pt Pet) (Pet, error) {
	if err := p.DB.Create(pt).Error; err != nil {
		return pt, err
	}
	return pt, nil
}

func (p *PetRepository) Get(id string) (Pet, error) {
	var r Pet

	if err := p.DB.Where("id = ?", id).First(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

func (p *PetRepository) Update(id string, pt Pet) (Pet, error) {
	if err := p.DB.Table("pets").Where("id = ?", id).Updates(pt).Error; err != nil {
		return pt, err
	}
	return pt, nil
}

func (p *PetRepository) Delete(id string) (Pet, error) {
	var r Pet
	if err := p.DB.Table("pets").Where("id = ?", id).Delete(&r).Error; err != nil {
		return r, err
	}
	return r, nil

}
