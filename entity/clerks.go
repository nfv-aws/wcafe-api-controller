package entity

type Clerk struct {

	// 店員の識別子
	Id string `dynamo:"id"`

	// 氏名
	Name string `dynamo:"name" binding:"required"`
}

type Clerks struct {
	Clerks *[]Clerk `dynamo:"clerks"`
}
