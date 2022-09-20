// Package services 面向服务
package services

import (
	"LuckyGo/dao"
	"LuckyGo/model"
)

type CodeService interface {
	GetAll() []model.Gifts
	CountAll() int64
	Get(id int) *model.Gifts
	Delete(id int) error
	Update(data *model.Gifts, columns []string) error
	Create(data *model.Gifts) error
}
type codeService struct {
	dao *dao.GiftDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *codeService) GetAll() []model.Gifts {
	return s.dao.GetAll()
}
func (s *codeService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *codeService) Get(id int) *model.Gifts {
	return s.dao.Get(id)
}
func (s *codeService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *codeService) Update(data *model.Gifts, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *codeService) Create(data *model.Gifts) error {
	return s.dao.Create(data)
}
