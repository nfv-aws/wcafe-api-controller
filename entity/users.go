package entity

import (
	"time"
)

type User struct {

	// 商品の識別子
	Id string `gorm:"type:varchar(255)" json:"id"`

	// 商品種
	Number int32 `gorm:"type:int(32) AUTO_INCREMENT;NOT NULL;unique" json:"number"`

	// 商品名
	Name string `gorm:"type:varchar(255)" json:"name"`

	// 商品の年齢
	Address string `gorm:"type:varchar(255)" json:"address"`

	// 店の識別子
	Email string `gorm:"type:varchar(255)" json:"email"`

	//作成日時
	CreatedAt time.Time `json:"created_time"`

	//更新日時
	UpdatedAt time.Time `json:"updated_time"`
}

type Users struct {
	Users *[]User `json:"users"`
}
