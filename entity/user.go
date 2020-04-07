package entity

type User struct {

	// 商品の識別子
	Id string `json:"id"`

	// 商品種
	Species string `json:"species"`

	// 商品名
	Name string `json:"name"`

	// 商品の年齢
	Age int32 `json:"age"`

	// 店の識別子
	StoreId string `json:"store_id"`
}

type Users struct {
	Users *[]User
}
