package controllers

import (
	"fmt"
	"myblog/models"
)

type MainController struct {
	BaseController
}

func (m *MainController) Get() {
	tag := m.GetString("tag")
	page, _ := m.GetInt("page")
	var artList []models.Article
	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		m.Data["HasFooter"] = false
	} else {
		if page < 0 {
			page = 1
		}
		//设置分页
		artList, _= models.FindArticleWithPage(page)
		fmt.Println(artList)
		//if error !=nil {
			//m.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
		//}
		//m.Data['']
	}
	m.Data["Content"] = models.MakeHomeBlocks(artList, m.IsLogin)
	m.TplName = "home.html"
	fmt.Println(tag, page)
	return
}
