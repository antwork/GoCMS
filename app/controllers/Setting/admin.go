package controllers

// 管理员设置
import "strconv"
import "github.com/robfig/revel"
import "admin/app/models"

type Admin struct {
	*revel.Controller
}

//首页
func (c Admin) Index(admin *models.Admin) revel.Result {
	title := "管理员管理--GoCMS管理系统"

	admin_list := admin.GetByAll(0)

	c.Render(title, admin_list)
	return c.RenderTemplate("Setting/Admin/Index.html")
}

//添加管理员
func (c Admin) Add(admin *models.Admin) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加管理员--GoCMS管理系统"

		role := new(models.Role)
		role_list := role.GetByAll()

		c.Render(title, role_list)
		return c.RenderTemplate("Setting/Admin/Add.html")
	} else {

		var username string = c.Params.Get("username")
		if len(username) > 0 {
			admin.Username = username
		} else {
			c.Flash.Error("请输入用户名!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var password string = c.Params.Get("password")
		if len(password) > 0 {
			admin.Password = password
		} else {
			c.Flash.Error("请输入密码!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var pwdconfirm string = c.Params.Get("pwdconfirm")
		if len(pwdconfirm) > 0 {
			if password != pwdconfirm {
				c.Flash.Error("两次输入密码不一致!")
				c.Flash.Out["url"] = "/Admin/Add/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("请输入确认密码!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var email string = c.Params.Get("email")
		if len(email) > 0 {
			admin.Email = email
		} else {
			c.Flash.Error("请输入E-mail!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var realname string = c.Params.Get("realname")
		if len(realname) > 0 {
			admin.Realname = realname
		} else {
			c.Flash.Error("请输入真实姓名!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var lang string = c.Params.Get("lang")
		if len(lang) > 0 {
			admin.Lang = lang
		} else {
			c.Flash.Error("请选择语言!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var roleid string = c.Params.Get("roleid")
		if len(roleid) > 0 {

			Roleid, err := strconv.ParseInt(roleid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin.Roleid = Roleid
		} else {
			c.Flash.Error("请选择所属角色!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		if admin.Save() {
			c.Flash.Success("添加管理员成功!")
			c.Flash.Out["url"] = "/Admin/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加管理员失败!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑管理员
func (c Admin) Edit(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "编辑管理员--GoCMS管理系统"

		role := new(models.Role)
		role_list := role.GetByAll()

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(Id)

			c.Render(title, admin_info, role_list)
		} else {
			c.Render(title, role_list)
		}

		return c.RenderTemplate("Setting/Admin/Edit.html")
	} else {

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			var username string = c.Params.Get("username")
			if len(username) > 0 {
				admin.Username = username
			} else {
				c.Flash.Error("请输入用户名!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var password string = c.Params.Get("password")
			if len(password) > 0 {
				admin.Password = password
			}

			var pwdconfirm string = c.Params.Get("pwdconfirm")
			if len(pwdconfirm) > 0 {
				if password != pwdconfirm {
					c.Flash.Error("两次输入密码不一致!")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}
			}

			var email string = c.Params.Get("email")
			if len(email) > 0 {
				admin.Email = email
			} else {
				c.Flash.Error("请输入E-mail!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var realname string = c.Params.Get("realname")
			if len(realname) > 0 {
				admin.Realname = realname
			} else {
				c.Flash.Error("请输入真实姓名!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var lang string = c.Params.Get("lang")
			if len(lang) > 0 {
				admin.Lang = lang
			} else {
				c.Flash.Error("请选择语言!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var roleid string = c.Params.Get("roleid")
			if len(roleid) > 0 {

				Roleid, err := strconv.ParseInt(roleid, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin.Roleid = Roleid
			} else {
				c.Flash.Error("请选择所属角色!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if admin.Edit(Id) {
				c.Flash.Success("编辑管理员成功!")
				c.Flash.Out["url"] = "/Admin/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑管理员失败!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑管理员失败!")
			c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}

	}
}

//删除管理员
func (c Admin) Delete(admin *models.Admin) revel.Result {
	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if admin.DelByID(Id) {
			data["status"] = "1"
			data["message"] = "删除成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "删除失败!"
			return c.RenderJson(data)
		}
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
		return c.RenderJson(data)
	}
}
