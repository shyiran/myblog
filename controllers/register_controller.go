package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"myblog/common"
	"myblog/models"
	"time"
)

type RegisterController struct {
	web.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

//处理注册
func (r *RegisterController) Post() {
	username := r.GetString("username:")
	password := r.GetString("password")
	repassword := r.GetString("repassword")
	fmt.Println(username+"sssssssssssss")
	fmt.Println("ddddddddddddddddddddddddddddddddddddddddd")
	fmt.Println(username, "ddvdvdvv",password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)

	fmt.Println("id:", id)
	if id > 0 {
		r.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		r.ServeJSON()
		return
	}
	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = common.MD5(password)
	fmt.Println("md5后：", password)
	//fmt.Println("md5后：", password)
	user := models.User{0, "uuuuuu", password, "276657548@qq.com",0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		r.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		r.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	r.ServeJSON()
}