package entity

type Store struct {

	// 店の識別子
	Id string `json:"id"; not null; primary key; type:varchar(255)`

	// 店名
	Name string `json:"name"; not null; type:varchar(255)`

	// 店の種類
	Tag string `json:"tag"; not null; type:varchar(255)`

	// 住所
	Address string `json:"address"; not null; type:varchar(255)`

	// 店の強み
	StrongPoint string `json:"strong_point,omitempty"; type:varchar(255)`
}

type Stores struct {
	Stores *[]Store `json:"stores"`
}
