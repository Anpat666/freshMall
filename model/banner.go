package model

type Banner struct {
	BannerId    string `json:"bannerId" gorm:"column:banner_id"`
	Url         string `json:"url" gorm:"column:url"`
	RedirectUrl string `json:"redirectUrl" gorm:"column:redirect_url"`
	Order       int    `json:"order" gorm:"column:order"`
	CreateUser  string `json:"createUser" gorm:"column:create_user"`
	UpdataUser  string `json:"updataUser" gorm:"column:updata_user"`
}
