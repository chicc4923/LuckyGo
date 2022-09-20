package main

import (
	"LuckyGo/bootrap"
	"LuckyGo/web/routes"
	"fmt"
)

var port = 8080

func newApp() *bootrap.Bootstrapper {
	//初始化应用
	app := bootrap.New("Go抽奖系统", "chicc")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
