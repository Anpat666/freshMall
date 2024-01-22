package repository

import (
	"errors"
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
	Delete(Product model.Product) (bool, error)
}

func (repo *ProductRepository) List(req *query.ListQuery) (products []*model.Product, err error) {
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page)
	sort := utils.Sort(req.Sort)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(int(limit)).Offset(int(offset)).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var products []model.Product
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&products).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *ProductRepository) Get(Product model.Product) (model.Product, error) {
	if err := repo.DB.Where(&Product).Find(&Product).Error; err != nil {
		return Product, err
	}
	return Product, nil
}

func (repo *ProductRepository) Exist(Product model.Product) *model.Product {
	if Product.ProductName != "" {
		var temp model.Product
		repo.DB.Where("product_name=?", Product.ProductName).First(&temp)
		return &temp
	}
	return nil
}

func (repo *ProductRepository) ExistByProductId(id string) *model.Product {
	var Product model.Product
	repo.DB.Where("product_id=?").Find(&Product)
	return &Product
}

func (repo *ProductRepository) Add(Product model.Product) (*model.Product, error) {
	if exist := repo.Exist(Product); exist != nil && exist.ProductName != "" {
		return nil, errors.New("商品已存在")
	}
	err := repo.DB.Create(&Product).Error
	if err != nil {
		return nil, errors.New("商品添加失败")
	}
	return &Product, nil
}

func (repo *ProductRepository) Edit(Product model.Product) (bool, error) {
	if Product.ProductId == "" {
		return false, errors.New("请输入正确的商品ID")
	}
	p := &model.Product{}
	err := repo.DB.Model(p).Where("product_id=?", Product.ProductId).Updates(map[string]interface{}{
		"product_name":           Product.ProductName,
		"product_intro":          Product.ProductIntro,
		"category_id":            Product.CategoryId,
		"product_cover_img":      Product.ProductCoverImg,
		"product_banner":         Product.ProductBanner,
		"original_price":         Product.OriginalPrice,
		"selling_price":          Product.SellingPrice,
		"stock_num":              Product.StockNum,
		"tag":                    Product.Tag,
		"sell_status":            Product.SellStatus,
		"product_detail_content": Product.ProductDatailContent,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *ProductRepository) Delete(Product model.Product) (bool, error) {
	err := repo.DB.Model(&Product).Where("product_id=?", Product.ProductId).Update("is_delete", Product.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
