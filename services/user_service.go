// Package services 面向服务
package services

import (
	"LuckyGo/dao"
	"LuckyGo/model"
)

type UserService interface {
	GetAll() []model.Gifts
	CountAll() int64
	Get(id int) *model.Gifts
	Delete(id int) error
	Update(data *model.Gifts, columns []string) error
	Create(data *model.Gifts) error
}
type userService struct {
	dao *dao.GiftDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *userService) GetAll() []model.Gifts {
	return s.dao.GetAll()
}
func (s *userService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *userService) Get(id int) *model.Gifts {
	return s.dao.Get(id)
}
func (s *userService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *userService) Update(data *model.Gifts, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *userService) Create(data *model.Gifts) error {
	return s.dao.Create(data)
}
