package service

import (
	"errors"
	"freshMall/model"
	"freshMall/repository"
)

type UserSrv interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}

func (u UserService) List(req *query.ListQuery) (users []*model.User, err error) {
	return u.Repo.List(req)
}

func (u UserService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return u.Repo.GetTotal(req)
}
func (u UserService) Get(user model.User) (*model.User, error) {
	return u.Repo.Get(user)
}
func (u UserService) Exist(user model.User) *model.User {
	return u.Repo.Exist(user)
}
func (u UserService) Add(user model.User) (*model.User, error) {
	result := u.Repo.ExistByMobile(user.Mobile)
	if result != nil {
		return nil, errors.New("用户已存在")
	}
	user.UserId = uuid.New().String()
	return u.Repo.Add(user)
}
func (u UserService) Edit(user model.User) (bool, error) {
	return u.Repo.Edit(user)
}
func (u UserService) Delete(id string) (bool, error) {
	return u.Repo.Delete(id)
}
