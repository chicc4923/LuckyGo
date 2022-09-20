package dao

import (
	"LuckyGo/model"
	"github.com/go-xorm/xorm"
	"log"
)

type blackipdao struct {
	engine *xorm.Engine
}

func NewBlackIpDao(engine *xorm.Engine) *blackipdao {
	return &blackipdao{
		engine: engine,
	}
}
func (d *blackipdao) Get(id int) *model.Gifts {
	data := &model.Gifts{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *blackipdao) GetAll() []model.Gifts {
	datalist := make([]model.Gifts, 0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)
	if err != nil {
		log.Println("gift_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}
func (d *blackipdao) CountAll() int64 {
	num, err := d.engine.Count(&model.Gifts{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *blackipdao) Delete(id int) error {
	//软删除，修改数据状态
	data := &model.Gifts{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *blackipdao) Updata(data *model.Gifts, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *blackipdao) Create(data *model.Gifts) error {
	_, err := d.engine.Insert(data)
	return err
}
func (d *blackipdao) GetByIp(ip string) *model.Blackip {
	datalist := make([]model.Blackip, 0)
	err := d.engine.Where("ip=?", ip).
		Desc("id").
		Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	} else {
		return &datalist[0]
	}
}
