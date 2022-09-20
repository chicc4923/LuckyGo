package model

type Gifts struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Title        string `xorm:"VARCHAR(255)"`
	PrizeNum     int    `xorm:"INT(11)"`
	LeftNum      int    `xorm:"INT(11)"`
	PrizeCode    string `xorm:"VARCHAR(50)"`
	PrizeTime    int    `xorm:"INT(11)"`
	Img          string `xorm:"VARCHAR(255)"`
	Displayorder int    `xorm:"INT(11)"`
	Gtype        int    `xorm:"INT(11)"`
	Gdata        string `xorm:"VARCHAR(255)"`
	TimeBegin    int    `xorm:"INT(11)"`
	TimeEnd      int    `xorm:"INT(11)"`
	PrizeData    string `xorm:"MEDIUMTEXT"`
	PrizeBegin   int    `xorm:"INT(11)"`
	PrizeEnd     int    `xorm:"INT(11)"`
	SysStatus    int    `xorm:"SMALLINT(6)"`
	SysCreated   int    `xorm:"INT(11)"`
	SysUpdated   int    `xorm:"INT(11)"`
	SysIp        string `xorm:"VARCHAR(50)"`
}
