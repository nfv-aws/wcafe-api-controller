package entity

type Store struct {

	// 店の識別子
	Id string `gorm:"NOT NULL;primary_key;type:varchar(255)" json:"id"`

	// 店名
	Name string `gorm:"NOT NULL;type:varchar(255)" json:"name"`

	// 店の種類
	Tag string `gorm:"NOT NULL;type:varchar(255)" json:"tag"`

	// 住所
	Address string `gorm:"NOT NULL;type:varchar(255)" json:"address"`

	// 店の強み
	StrongPoint string `gorm:"type:varchar(255)" json:"strong_point"`
}

type Stores struct {
	Stores *[]Store `json:"stores"`
}
