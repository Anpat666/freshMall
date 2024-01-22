package model

type Category struct {
	CategoryId string `json:"categoryId" gorm:"column:category_id"`
	Name       string `json:"name" gorm:"column:name"`
	Desc       string `json:"desc" gorm:"column:desc"`
	Order      string `json:"order" gorm:"column:order"`
	ParentId   string `json:"parent_id" gorm:"column:parent_id"`
	IsDeleted  bool   `json:"is_deleted" gorm:"column:is_deleted"`
}

type CategoryResult struct {
	C1CategoryId string `gorm:"c1_category_id"`
	C1Name       string `gorm:"column:c1_name"`
	C1Desc       string `gorm:"column:c1_desc"`
	C1Order      int    `gorm:"column:c1_order"`
	C1ParentId   string `gorm:"column:c1_parent_id"`

	C2CategoryId string `gorm:"c2_category_id"`
	C2Name       string `gorm:"column:c2_name"`
	C2Order      int    `gorm:"column:c2_order"`
	C2ParentId   string `gorm:"column:c2_parent_id"`

	C3CategoryId string `gorm:"c3_category_id"`
	C3Name       string `gorm:"column:c3_name"`
	C3Order      int    `gorm:"column:c3_order"`
	C3ParentId   string `gorm:"column:c3_parent_id"`
}
