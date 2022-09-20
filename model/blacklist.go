package model

type Blacklist struct {
	Id int `xorm:"not null pk autoincr INT(11)"`
}
