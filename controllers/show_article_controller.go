package controllers

import (
	"fmt"
	"strconv"
	"myblog/models"
	"myblog/common"
)

type ShowArticleController struct {
	BaseController
}
func (s *ShowArticleController) Get() {

	idStr := s.Ctx.Input.Param(":id")

	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)

	s.Data["Title"] = art.Title
	//this.Data["Content"] = art.Content
	s.Data["Content"] = common.SwitchMarkdownToHtml(art.Content)
	s.TplName = "show_article.html"
}


