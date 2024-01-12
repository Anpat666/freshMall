package repository

import (
	"freshMall/model"
	"freshMall/utils"

	"github.com/aws/aws-sdk-go/private/protocol/query"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	List(req *query.ListQuery) (Products []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Product model.Product) (*model.Product, error)
	Exist(Product model.Product) *model.Product
	ExistByProductId(id string) *model.Product
	Add(Product model.Product) (*model.Product, error)
	Edit(Product model.Product) (bool, error)
	Delete(p model.Product) (bool, error)
}

func (repo *ProductRepository) List(req *query.ListQuery) (products []*model.Product, err error) {
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page)
	sort := utils.Sort(req.Sort)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
