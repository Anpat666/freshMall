package repository

import (
	"errors"
	"freshMall/model"
	"freshMall/utils"

	"github.com/aws/aws-sdk-go/private/protocol/query"
	"gorm.io/gorm"
)

type BannerRepository struct {
	DB *gorm.DB
}

type BannerRepoInterface interface {
	List(req *query.ListQuery) (banner []*model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Banner model.Banner) (*model.Banner, error)
	Exist(Banner model.Banner) *model.Banner
	ExistByBannerID(id string) *model.Banner
	Add(Banner model.Banner) (*model.Banner, error)
	Edit(Banner model.Banner) (bool, error)
	Delete(Banner model.Banner) (bool, error)
}

func (b *BannerRepository) List(req *query.ListQuery) (banners []*model.Banner, err error) {
	db := b.DB
	limit, offset := utils.Page(req.limit, req.offset)
	sort := utils.Sort(req.sort)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(int(limit)).Offset(int(offset)).Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}

func (b *BannerRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	db := b.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	var banners []model.Banner
	if err := db.Find(banners).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (b *BannerRepository) Get(Banner model.Banner) (*model.Banner, error) {
	if err := b.DB.Model(&Banner).Where(Banner).Find(&Banner).Error; err != nil {
		return nil, err
	}
	return &Banner, nil
}
func (b *BannerRepository) Exist(Banner model.Banner) *model.Banner {
	if Banner.Url != "" && Banner.RedirectUrl != "" {
		b.DB.Model(&Banner).Where("url=? and redirect_url=?", Banner.Url, Banner.RedirectUrl).First(&Banner)
		return &Banner
	}
	return nil
}

func (b *BannerRepository) ExistByBannerID(id string) *model.Banner {
	var banner model.Banner
	b.DB.Where("banner_id=?", id).First(&banner)
	return &banner
}

func (b *BannerRepository) Add(Banner model.Banner) (*model.Banner, error) {
	exist := b.Exist(Banner)
	if exist != nil && exist.Url == Banner.Url && exist.RedirectUrl == Banner.RedirectUrl {
		return nil, errors.New("轮播图已经存在")
	}
	err := b.DB.Create(&Banner).Error
	if err != nil {
		return nil, errors.New("轮播图添加失败")
	}
	return &Banner, nil
}

func (b *BannerRepository) Edit(Banner model.Banner) (bool, error) {
	if Banner.BannerId == "" {
		return false, errors.New("参数错误")
	}
	if err := b.DB.Model(&Banner).Where("banner_id=?", Banner.BannerId).Updates(map[string]interface{}{
		"banner_url":   Banner.Url,
		"redirect_url": Banner.RedirectUrl,
		"order":        Banner.Order,
	}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (b *BannerRepository) Delete(Banner model.Banner) (bool, error) {
	err := b.DB.Model(Banner).Where("banner_id=?", Banner.BannerId).Delete(&Banner).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
