package entity

type Clerk struct {

	// 店員の識別子
	NameId string `dynamo:"name_id"`

	// 氏名
	Name string `dynamo:"name"`
}

type Clerks struct {
	Clerks *[]Clerk `dynamo:"clerks"`
}
