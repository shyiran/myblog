package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"myblog/common"
	"myblog/models"
)

type LoginController struct {
	web.Controller

}
//登录
func (l *LoginController) Get() {
	l.TplName = "login.html"
}
func (l *LoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")

	//fmt.Println("username:", username, ",password:", password)
	id := models.QueryUserWithParam(username, common.MD5(password))

	//id_s:=l.GetString("id")
	//fmt.Println("SSSSSSSSSSSSSSSSSSSSSssssssss")
	//fmt.Println("id:", id)
	//fmt.Println(id)
	if id > 0 {
		/*
			设置了session后悔将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
			因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		*/
		l.SetSession("loginuser", username)
		l.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		l.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	l.ServeJSON()
}

