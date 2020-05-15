package entity

import (
	"time"
)

type User struct {

	// ユーザの識別子
	Id string `gorm:"type:varchar(255)" json:"id"`

	// 会員番号
	Number int32 `gorm:"type:int(32) AUTO_INCREMENT;NOT NULL;unique" json:"number"  binding:"required"`

	// 氏名
	Name string `gorm:"type:varchar(255)" json:"name"`

	// 住所
	Address string `gorm:"type:varchar(255)" json:"address"`

	// メールアドレス
	Email string `gorm:"type:varchar(255)" json:"email"`

	//作成日時
	CreatedAt time.Time `json:"created_time"`

	//更新日時
	UpdatedAt time.Time `json:"updated_time"`
}

type Users struct {
	Users *[]User `json:"users"`
}
