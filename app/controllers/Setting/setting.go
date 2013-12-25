package controllers

//设置首页

import "github.com/robfig/revel"

type Setting struct {
	*revel.Controller
}

func (c Setting) Index() revel.Result {
	title := "设置--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Setting/Index.html")
}
