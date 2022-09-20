// Package services 面向服务
package services

import (
	"LuckyGo/dao"
	"LuckyGo/model"
)

type UserDayService interface {
	GetAll() []model.Gifts
	CountAll() int64
	Get(id int) *model.Gifts
	Delete(id int) error
	Update(data *model.Gifts, columns []string) error
	Create(data *model.Gifts) error
}
type userdayService struct {
	dao *dao.GiftDao
}

func NewUserDayService() UserDayService {
	return &userdayService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *userdayService) GetAll() []model.Gifts {
	return s.dao.GetAll()
}
func (s *userdayService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *userdayService) Get(id int) *model.Gifts {
	return s.dao.Get(id)
}
func (s *userdayService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *userdayService) Update(data *model.Gifts, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *userdayService) Create(data *model.Gifts) error {
	return s.dao.Create(data)
}
