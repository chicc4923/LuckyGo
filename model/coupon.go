package model

type Coupon struct {
	Id int `xorm:"not null pk autoincr INT(11)"`
}
