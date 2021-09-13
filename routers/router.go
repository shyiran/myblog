package routers

import (
	"github.com/beego/beego/v2/server/web"
	"myblog/controllers"
)

func init() {
    web.Router("/", &controllers.MainController{})
	//注册
	web.Router("/register", &controllers.RegisterController{})
	//登录
	web.Router("/login", &controllers.LoginController{})
	//关于我页面
	web.Router("/aboutme",&controllers.AboutMeController{})
	//相册
	web.Router("/album", &controllers.AlbumController{})
	//标签
	web.Router("/tags", &controllers.TagsController{})
	//文件上传
	web.Router("/upload", &controllers.UploadController{})
	//显示文章详情
	web.Router("/article/:id", &controllers.ShowArticleController{})
}