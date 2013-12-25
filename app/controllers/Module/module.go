package controllers

//模块首页
import "github.com/robfig/revel"

type Module struct {
	*revel.Controller
}

func (c Module) Index() revel.Result {
	title := "模块--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Module/Index.html")
}
