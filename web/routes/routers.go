package routes

import (
	"LuckyGo/bootrap"
	"LuckyGo/services"
	"LuckyGo/web/controllors"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(b *bootrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userdayService := services.NewUserDayService()
	blackipService := services.NewBlackIpService()

	index := mvc.New(b.Party("/"))
	index.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	index.Handle(new(controllors.IndexControllor))

}
