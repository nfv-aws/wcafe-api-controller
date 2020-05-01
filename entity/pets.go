package entity

import (
	"time"
)

type Pet struct {

	// 商品の識別子
	Id string `gorm:"type:varchar(255)" json:"id"`

	// 商品種
	Species string `gorm:"NOT NULL;type:varchar(255)" json:"species"`

	// 商品名
	Name string `gorm:"NOT NULL;type:varchar(255)" json:"name"`

	// 商品の年齢
	Age int32 `gorm:"NOT NULL;type:int(11)" json:"age"`

	// 店の識別子
	StoreId string `gorm:"NOT NULL;type:varchar(255)" json:"store_id"`

	//作成日時
	CreatedAt time.Time `json:"created_time"`

	//更新日時
	UpdatedAt time.Time `json:"updated_time"`

	// ステータス
	Status string `json:"status"`
}

type Pets struct {
	Pets *[]Pet `json:"pets"`
}
