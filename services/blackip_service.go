// Package services 面向服务
package services

import (
	"LuckyGo/dao"
	"LuckyGo/model"
)

type BlackIpService interface {
	GetAll() []model.Gifts
	CountAll() int64
	Get(id int) *model.Gifts
	Delete(id int) error
	Update(data *model.Gifts, columns []string) error
	Create(data *model.Gifts) error
}
type blackIpService struct {
	dao *dao.GiftDao
}

func NewBlackIpService() BlackIpService {
	return &blackIpService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *blackIpService) GetAll() []model.Gifts {
	return s.dao.GetAll()
}
func (s *blackIpService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *blackIpService) Get(id int) *model.Gifts {
	return s.dao.Get(id)
}
func (s *blackIpService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *blackIpService) Update(data *model.Gifts, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *blackIpService) Create(data *model.Gifts) error {
	return s.dao.Create(data)
}
