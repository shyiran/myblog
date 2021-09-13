package controllers

import "myblog/models"

type AlbumController struct {
	BaseController
}
func (a *AlbumController) Get(){
	albums,err:=models.FindAllAlbums()
	if err!=nil{
		err.Error()
	}
	a.Data["Album"] = albums
	a.TplName = "album.html"
}