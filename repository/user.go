package repository

import (
	"errors"
	"freshMall/model"
	"freshMall/utils"

	"github.com/aws/aws-sdk-go/private/protocol/query"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *query.ListQuery) (user []model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	ExistByMobile(mobile string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(u model.User) (bool, error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users []model.User, err error) {
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order("id desc").Limit(int(limit)).Offset(int(offset)).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var users []model.User
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *UserRepository) Get(user model.User) (*model.User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Exist(user model.User) *model.User {
	var count int64
	repo.DB.Find(&user).Where("nick_name=?", user.NickName).Count(&count)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var user model.User
	repo.DB.Where("user_id=?", id).Find(&user)
	return &user
}

func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist != nil {
		return nil, errors.New("用户已存在")
	}
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, errors.New("用户注册失败")
	}
	return &user, nil
}

func (repo *UserRepository) Edit(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserId).Updates(map[string]interface{}{
		"nick_name": user.NickName,
		"mobile":    user.Mobile,
		"address":   user.Address,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) Delete(u model.User) (bool, error) {
	err := repo.DB.Model(&u).Where("user_id=?", u.UserId).Update("is_deleted", u.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
