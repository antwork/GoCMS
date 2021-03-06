package controllers

// 菜单设置
import "strconv"
import "github.com/robfig/revel"
import "admin/app/models"

type Menu struct {
	*revel.Controller
}

//首页
func (c Menu) Index(menu *models.Menu) revel.Result {
	title := "菜单管理--GoCMS管理系统"

	menus := menu.GetMenuHtml()

	c.Render(title, menus)
	return c.RenderTemplate("Setting/Menu/Index.html")
}

//添加菜单
func (c Menu) Add(menu *models.Menu) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加菜单--GoCMS管理系统"

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//返回菜单Option的HTML
			menus := menu.GetMenuOptionHtml(Id)

			c.Render(title, menus, Id)
		} else {
			//返回菜单Option的HTML
			menus := menu.GetMenuOptionHtml(0)
			c.Render(title, menus)
		}

		return c.RenderTemplate("Setting/Menu/Add.html")
	} else {

		var pid string = c.Params.Get("pid")
		if len(pid) > 0 {
			Pid, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			menu.Pid = Pid
		} else {
			c.Flash.Error("请选择父菜单!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var name string = c.Params.Get("name")
		if len(name) > 0 {
			menu.Name = name
		} else {
			c.Flash.Error("请输入菜单名称!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var url string = c.Params.Get("url")
		if len(url) > 0 {
			menu.Url = url
		} else {
			c.Flash.Error("请输入菜单地址!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var order string = c.Params.Get("order")
		if len(order) > 0 {
			Order, err := strconv.ParseInt(order, 10, 16)
			if err != nil {
				revel.WARN.Println(err)
			}
			menu.Order = Order
		} else {
			c.Flash.Error("请输入排序!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var data string = c.Params.Get("data")
		menu.Data = data

		var display string = c.Params.Get("display")
		if len(display) > 0 {
			Display, err := strconv.ParseInt(display, 10, 8)
			if err != nil {
				revel.WARN.Println(err)
			}
			menu.Display = Display
		} else {
			c.Flash.Error("请选择是否显示菜单!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		if menu.Save() {
			c.Flash.Success("添加菜单成功")
			c.Flash.Out["url"] = "/Menu/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加菜单失败")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑菜单
func (c Menu) Edit(menu *models.Menu) revel.Result {
	if c.Request.Method == "GET" {
		title := "编辑菜单--GoCMS管理系统"

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//获取菜单信息
			menu_info := menu.GetById(Id)

			//返回菜单Option的HTML
			menus := menu.GetMenuOptionHtml(menu_info.Pid)

			c.Render(title, menus, menu_info)
		} else {

			//返回菜单Option的HTML
			menus := menu.GetMenuOptionHtml(0)

			c.Render(title, menus)
		}
		return c.RenderTemplate("Setting/Menu/Edit.html")
	} else {

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			var pid string = c.Params.Get("pid")
			if len(pid) > 0 {
				Pid, err := strconv.ParseInt(pid, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				menu.Pid = Pid
			} else {
				c.Flash.Error("请选择父菜单!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var name string = c.Params.Get("name")
			if len(name) > 0 {
				menu.Name = name
			} else {
				c.Flash.Error("请输入菜单名称!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var url string = c.Params.Get("url")
			if len(url) > 0 {
				menu.Url = url
			} else {
				c.Flash.Error("请输入菜单地址!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var order string = c.Params.Get("order")
			if len(order) > 0 {
				Order, err := strconv.ParseInt(order, 10, 16)
				if err != nil {
					revel.WARN.Println(err)
				}
				menu.Order = Order
			} else {
				c.Flash.Error("请输入排序!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var data string = c.Params.Get("data")
			menu.Data = data

			var display string = c.Params.Get("display")
			if len(display) > 0 {
				Display, err := strconv.ParseInt(display, 10, 8)
				if err != nil {
					revel.WARN.Println(err)
				}
				menu.Display = Display
			} else {
				c.Flash.Error("请选择是否显示菜单!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if menu.Edit(Id) {
				c.Flash.Success("编辑菜单成功")
				c.Flash.Out["url"] = "/Menu/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑菜单失败")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑菜单失败")
			c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	}
}

//删除
func (c Menu) Delete(menu *models.Menu) revel.Result {
	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) <= 0 {
		data["status"] = "0"
		data["message"] = "参数错误!"
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	if menu.DelByID(Id) {
		data["status"] = "1"
		data["message"] = "删除成功!"
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
	}

	return c.RenderJson(data)
}
