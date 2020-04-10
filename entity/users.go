package entity

type User struct {

	// 商品の識別子
	Id string `json:"id"`

	// 商品種
	Number int32 `json:"number"`

	// 商品名
	Name string `json:"name"`

	// 商品の年齢
	Address string `json:"address"`

	// 店の識別子
	Email string `json:"email"`
}

type Users struct {
	Users *[]User `json:"users"`
}
