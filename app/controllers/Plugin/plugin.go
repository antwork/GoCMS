package controllers

//应用首页
import "github.com/robfig/revel"

type Plugin struct {
	*revel.Controller
}

func (c Plugin) Index() revel.Result {
	title := "设置--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Plugin/Index.html")
}
