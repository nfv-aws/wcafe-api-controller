package entity

type User struct {

	// 商品の識別子
	Id string `gorm:"NOT NULL;primary_key;type:varchar(255)" json:"id"`

	// 商品種
	Number int32 `gorm:"type:int(32)" json:"number"`

	// 商品名
	Name string `gorm:"type:varchar(255)" json:"name"`

	// 商品の年齢
	Address string `gorm:"type:varchar(255)" json:"address"`

	// 店の識別子
	Email string `gorm:"type:varchar(255)" json:"email"`
}

type Users struct {
	Users *[]User `json:"users"`
}
