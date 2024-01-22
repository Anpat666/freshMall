package repository

import (
	"errors"
	"freshMall/model"
	"freshMall/utils"

	"github.com/aws/aws-sdk-go/private/protocol/query"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	List(req *query.ListQuery) (Order []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Order model.Order) (*model.Order, error)
	Exist(Order model.Order) *model.Order
	ExistByOrderID(id string) *model.Order
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(Order model.Order) (bool, error)
}

func (o *OrderRepository) List(req *query.ListQuery) (Order []*model.Order, err error) {
	db := o.DB
	limit, offset := utils.Page(req.limit, req.offset)
	sort := utils.Sort(req.sort)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(int(limit)).Offset(int(offset)).Find(&Order).Error; err != nil {
		return nil, err
	}
	return Order, nil
}

func (o *OrderRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var order []model.Order
	db := o.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&order).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (o *OrderRepository) Get(Order model.Order) (*model.Order, error) {
	if err := o.DB.Where(&Order).Find(&Order).Error; err != nil {
		return nil, err
	}
	return &Order, nil
}

func (o *OrderRepository) Exist(Order model.Order) *model.Order {
	if Order.OrderId != "" {
		o.DB.Model(&Order).Where("order_id=?", Order.OrderId).Find(&Order)
		return &Order
	}
	return nil
}

func (o *OrderRepository) ExistByOrderID(id string) *model.Order {
	var order *model.Order
	o.DB.Where("order_id=?", id).First(&order)
	return order
}

func (o *OrderRepository) Add(Order model.Order) (*model.Order, error) {
	if err := o.DB.Create(&Order).Error; err != nil {
		return nil, err
	}
	return &Order, nil
}

func (o *OrderRepository) Edit(Order model.Order) (bool, error) {
	if Order.OrderId == "" {
		return false, errors.New("传入ID错误")
	}
	err := o.DB.Model(&Order).Where("order_id=?", Order.OrderId).Updates(map[string]interface{}{
		"nick_name": Order.NickName,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *OrderRepository) Delete(Order model.Order) (bool, error) {
	err := o.DB.Model(&Order).Where("order_id", Order.OrderId).Update("is_delete", Order.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
