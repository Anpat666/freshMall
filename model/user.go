package model

import "time"

type User struct {
	UserId    string    `json:"userId" gorm:"column:user_id"`
	NickName  string    `json:"nickName" gorm:"column:nick_name"`
	Mobile    string    `json:"mobile" gorm:"column:mobile" binding:"required"`
	Password  string    `json:"password" gorm:"column:password"`
	Address   string    `json:"address" gorm:"column:address"`
	IsDeleted bool      `json:"isDeleted" gorm:"column:is_deleted"`
	CreateAt  time.Time `json:"createAt" gorm:"column:create_at"`
	UpdataAt  time.Time `json:"updataAt" gorm:"column:updata_at"`
}
