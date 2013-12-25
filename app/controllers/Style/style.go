package controllers

//界面首页
import "github.com/robfig/revel"

type Style struct {
	*revel.Controller
}

func (c Style) Index() revel.Result {
	title := "设置--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Style/Index.html")
}
