package controllers

//我的面板首页
import "github.com/robfig/revel"

type Panel struct {
	*revel.Controller
}

func (c Panel) Index() revel.Result {
	title := "设置--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Panel/Index.html")
}
