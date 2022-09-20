// Package services 面向服务
package services

import (
	"LuckyGo/dao"
	"LuckyGo/model"
)

type ResultService interface {
	GetAll() []model.Gifts
	CountAll() int64
	Get(id int) *model.Gifts
	Delete(id int) error
	Update(data *model.Gifts, columns []string) error
	Create(data *model.Gifts) error
}
type resultService struct {
	dao *dao.GiftDao
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *resultService) GetAll() []model.Gifts {
	return s.dao.GetAll()
}
func (s *resultService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *resultService) Get(id int) *model.Gifts {
	return s.dao.Get(id)
}
func (s *resultService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *resultService) Update(data *model.Gifts, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *resultService) Create(data *model.Gifts) error {
	return s.dao.Create(data)
}
