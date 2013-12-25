package controllers

//内容首页
import "github.com/robfig/revel"

type Content struct {
	*revel.Controller
}

func (c Content) Index() revel.Result {
	title := "内容--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Content/Index.html")
}
