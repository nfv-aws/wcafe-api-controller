package entity

type Pet struct {

	// 商品の識別子
	Id string `json:"id"; unique; not null; primary key; read only; type:varchar(255)`

	// 商品種
	Species string `json:"species"; not null; type:varchar(255)`

	// 商品名
	Name string `json:"name"; not null; type:varchar(255)`

	// 商品の年齢
	Age int32 `json:"age"; not null; type:int(11)`

	// 店の識別子
	StoreId string `json:"store_id"; not null; type:varchar(255)`
}

type Pets struct {
	Pets *[]Pet `json:"pets"`
}
