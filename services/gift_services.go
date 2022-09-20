// Package services 面向服务
package services

import (
	"LuckyGo/dao"
	"LuckyGo/model"
)

type GiftService interface {
	GetAll() []model.Gifts
	CountAll() int64
	Get(id int) *model.Gifts
	Delete(id int) error
	Update(data *model.Gifts, columns []string) error
	Create(data *model.Gifts) error
}
type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *giftService) GetAll() []model.Gifts {
	return s.dao.GetAll()
}
func (s *giftService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *giftService) Get(id int) *model.Gifts {
	return s.dao.Get(id)
}
func (s *giftService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *giftService) Update(data *model.Gifts, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *giftService) Create(data *model.Gifts) error {
	return s.dao.Create(data)
}
