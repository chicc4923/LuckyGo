package dao

import (
	"LuckyGo/model"
	"github.com/go-xorm/xorm"
	"log"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}
func (d *CodeDao) Get(id int) *model.Gifts {
	data := &model.Gifts{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *CodeDao) GetAll() []model.Gifts {
	datalist := make([]model.Gifts, 0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)
	if err != nil {
		log.Println("gift_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}
func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&model.Gifts{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *CodeDao) Delete(id int) error {
	data := &model.Gifts{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *CodeDao) Update(data *model.Gifts, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *model.Gifts) error {
	_, err := d.engine.Insert(data)
	return err
}
