package entity

type Store struct {

	// 店の識別子
	Id string `json:"id"`

	// 店名
	Name string `json:"name"`

	// 店の種類
	Tag string `json:"tag"`

	// 住所
	Address string `json:"address"`

	// 店の強み
	StrongPoint string `json:"strong_point,omitempty"`
}

type Stores struct {
	Stores *[]Store `json:"stores"`
}
