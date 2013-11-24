package controllers

//日志管理

import "github.com/robfig/revel"

type Logs struct {
	*revel.Controller
}

func (c Logs) Index() revel.Result {
	return c.Render()
}
