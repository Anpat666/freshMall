package service

import (
	"freshMall/model"
	"freshMall/repository"
)

type UserSrv interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user model.User) (*model.User, error)
	Exist(User model.User) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}
