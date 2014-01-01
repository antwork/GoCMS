package controllers

//后台登陆
import "fmt"
import "strconv"
import "github.com/robfig/revel"
import "admin/app/models"
import "admin/lib"

type User struct {
	*revel.Controller
}

//首页
func (c *User) Index() revel.Result {
	title := "用户首页--GoCMS管理系统"
	return c.Render(title)
}

//登陆
func (c *User) Login(admin *models.Admin) revel.Result {

	if c.Request.Method == "GET" {
		title := "登陆--GoCMS管理系统"
		return c.Render(title)
	} else {
		var username string = c.Params.Get("username")
		var password string = c.Params.Get("password")
		var verify string = c.Params.Get("verify")

		data := make(map[string]string)

		if len(username) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = "请填写用户名!"
			return c.RenderJson(data)
		}

		if len(password) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = "请填写密码!"
			return c.RenderJson(data)
		}

		if len(verify) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = "请填写验证码!"
			return c.RenderJson(data)
		}

		admin_info := admin.GetByName(username)

		if admin_info.Id <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = "用户名错误!"
		} else if username == admin_info.Username && lib.Md5(password) == admin_info.Password {
			c.Session["UserID"] = fmt.Sprintf("%d", admin_info.Id)

			c.Flash.Success("欢迎您 " + admin_info.Realname)
			c.Flash.Out["url"] = "/"

			data["status"] = "1"
			data["url"] = "/Message/"
			data["message"] = "登陆成功!"
		} else {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = "密码错误!"
		}

		return c.RenderJson(data)
	}
}

//退出登陆
func (c *User) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}

	c.Flash.Success("安全退出")
	c.Flash.Out["url"] = "/User/Login/"
	return c.Redirect("/Message/")
}

//个人信息
func (c *User) EditInfo(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "个人信息--GoCMS管理系统"

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)
			c.Render(title, admin_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("User/EditInfo.html")
	} else {

		var realname string = c.Params.Get("realname")
		if len(realname) > 0 {
			admin.Realname = realname
		} else {
			c.Flash.Error("请输入真实姓名!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		var email string = c.Params.Get("email")
		if len(email) > 0 {
			admin.Email = email
		} else {
			c.Flash.Error("请输入电子邮件!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		var lang string = c.Params.Get("lang")
		if len(lang) > 0 {
			admin.Lang = lang
		} else {
			c.Flash.Error("请选择语言!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			if admin.Edit(UserID) {
				c.Flash.Success("修改成功!")
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("修改失败!")
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("未登陆，请先登陆!")
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}
}

//修改密码
func (c *User) EditPwd(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "修改密码--GoCMS管理系统"

		c.Render(title)

		return c.RenderTemplate("User/EditPwd.html")
	} else {

		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)

			var old_password string = c.Params.Get("old_password")
			if len(old_password) > 0 {
				if admin_info.Password != lib.Md5(old_password) {
					c.Flash.Error("旧密码不正确!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				}
			} else {
				return c.Redirect("/User/EditPwd/")
			}

			var new_password string = c.Params.Get("new_password")
			if len(new_password) > 0 {

			} else {
				c.Flash.Error("新密码不能为空!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			var new_pwdconfirm string = c.Params.Get("new_pwdconfirm")
			if len(new_pwdconfirm) > 0 {
				if new_pwdconfirm != new_password {
					c.Flash.Error("两次输入密码入不一致!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				} else {
					admin.Password = new_pwdconfirm
				}
			} else {
				c.Flash.Error("请输入重复新密码!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			if admin.Edit(UserID) {
				c.Flash.Success("修改成功!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("修改失败!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("未登陆，请先登陆!")
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}
}

//左侧导航菜单
func (c *User) Left(menu *models.Menu) revel.Result {

	title := "左侧导航--GoCMS管理系统"

	var pid string = c.Params.Get("pid")

	if len(pid) > 0 {
		Pid, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		//获取左侧导航菜单
		left_menu := menu.GetLeftMenuHtml(Pid)

		c.Render(title, left_menu)
	} else {
		//获取左侧导航菜单
		//默认获取 我的面板
		left_menu := menu.GetLeftMenuHtml(1)

		c.Render(title, left_menu)
	}
	return c.RenderTemplate("Public/left.html")
}
