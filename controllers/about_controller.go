package controllers
type AboutMeController struct {
	BaseController
}
func (a *AboutMeController)Get(){
	a.Data["wechat"]="wodeweixinhao"
	a.Data["qq"]="wodeQQhaoma"
	a.Data["tel"]="wodedianhua"
	a.TplName="aboutme.html"
}
