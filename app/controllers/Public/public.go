package controllers

//后台公用处理
import "github.com/robfig/revel"
import "admin/app/models"

type Public struct {
	*revel.Controller
}

//后台地图
func (c *Public) Map(menu *models.Menu) revel.Result {

	//返回后台地图
	map_html := menu.GetMenuMap()

	c.Render(map_html)
	return c.RenderTemplate("Public/map.html")
}

//生成网站首页
func (c *Public) CreateHtml() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/createhtml.html")
}

//搜索
func (c *Public) Search() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/search.html")
}

//消息提示
func (c *Public) Message() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/message.html")
}
