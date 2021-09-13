package main

import (
	"github.com/beego/beego/v2/server/web"
	"myblog/common"
	_ "myblog/routers"
)

func main() {
	common.InitMysql()
	web.Run()
}
