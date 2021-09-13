package controllers

import (
	"fmt"
	"myblog/models"
)

type TagsController struct {
	BaseController
}
func (t *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))
	t.Data["Tags"] = models.HandleTagsListData(tags)
	t.TplName = "tags.html"
}