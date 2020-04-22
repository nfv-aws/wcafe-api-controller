package entity

type User struct {

	// 商品の識別子
	Id string `json:"id";gorm:"not null;primary_key;type:varchar(255)"`

	// 商品種
	Number int32 `json:"number";gorm:"type:int(32)"`

	// 商品名
	Name string `json:"name";gorm:"type:varchar(255)"`

	// 商品の年齢
	Address string `json:"address";gorm:"type:varchar(255)"`

	// 店の識別子
	Email string `json:"email";gorm:"type:varchar(255)"`
}

type Users struct {
	Users *[]User `json:"users"`
}
