package entity

type Supply struct {

	// ユーザの識別子
	Id string `dynamo:"id" json:"id"`

	// 氏名
	Name string `dynamo:"name" json:"name" binding:"required"`

	Price int `dynamo:"price" json:"price" binding:"required"`

	Type string `dynamo:"type" json:"type"`
}

type Supplies struct {
	Supplies *[]Supply `json:"supplies"`
}
