package repository

import (
	"errors"
	"freshMall/model"
	"freshMall/utils"

	"github.com/aws/aws-sdk-go/private/protocol/query"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	List(req *query.ListQuery) (Categories []*model.Category, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryId(id string) *model.Category
	Add(Category model.Category) (*model.Category, error)
	Edit(Category model.Category) (bool, error)
	Delete(Category model.Category) (bool, error)
}

func (repo *CategoryRepository) List(req *query.ListQuery) (Categories []*model.Category, err error) {
	db := repo.DB
	limit, offset := utils.Page(req.limit, req.Page)
	sort := utils.Sort(req.sort)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(int(limit)).Offset(int(offset)).Find(&Categories).Error; err != nil {
		return nil, err
	}
	return Categories, nil
}

func (repo *CategoryRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var category []*model.Category
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&category).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *CategoryRepository) Get(id string) ([]*model.CategoryResult, error) {
	var list []*model.CategoryResult
	err := repo.DB.Table("category_result").Where("c1_category_id=? OR c2_category_id =? OR c3_category_id=?", id, id, id).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *CategoryRepository) Exist(Category model.Category) *model.Category {
	var category model.Category
	if Category.Name != "" {
		repo.DB.Find(&category).Where("name=?", Category.Name)
		return &category
	}
	return nil
}

func (repo *CategoryRepository) ExistByCategoryId(id string) *model.Category {
	var category model.Category
	repo.DB.Where("category_id=?", id).First(&category)
	return &category
}

func (repo *CategoryRepository) Add(Category model.Category) (*model.Category, error) {

	err := repo.DB.Create(&Category).Error
	if err != nil {
		return nil, errors.New("分类创建失败")
	}
	return &Category, nil
}

func (repo *CategoryRepository) Edit(Category model.Category) (bool, error) {
	err := repo.DB.Model(&Category).Where("category_id=?", Category.CategoryId).Updates(map[string]interface{}{
		"name":      Category.Name,
		"order":     Category.Order,
		"parent_id": Category.ParentId,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *CategoryRepository) Delete(Category model.Category) (bool, error) {
	err := repo.DB.Model(&Category).Where("category_id=?", Category.CategoryId).Update("is_deleted", Category.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
