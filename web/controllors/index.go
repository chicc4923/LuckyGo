package controllors

import (
	comm "LuckyGo/coom"
	"LuckyGo/model"
	"LuckyGo/services"
	"fmt"
	"github.com/kataras/iris/v12"
)

type IndexControllor struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.CodeService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserDayService
	ServuceBlackIp services.BlackIpService
}

func (c *IndexControllor) Get() string {
	c.Ctx.Header("Content-Type", "text/h")

	return "Welcome to Golang抽奖系统，<a href='index.html'>开始抽奖</a>"
}

func (c *IndexControllor) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})

	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.ServiceGift.GetAll()
	list := make([]model.Gifts, 0)

	for _, data := range datalist {
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}

func (c *IndexControllor) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	return rs
}

func (c *IndexControllor) GetLogin() {
	uid := comm.Random(10000)
	loginuser := model.ObjLoginUser{
		Uid:      uid,
		UserName: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	comm.SetLoginUser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), "../public/index.html?from=login")
}

func (c *IndexControllor) GetLogout() {
	comm.SetLoginUser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), "./public/index.html?from=logout")
}
