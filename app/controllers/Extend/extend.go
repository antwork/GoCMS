package controllers

//扩展首页
import "github.com/robfig/revel"

type Extend struct {
	*revel.Controller
}

func (c Extend) Index() revel.Result {
	title := "扩展--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Extend/Index.html")
}
