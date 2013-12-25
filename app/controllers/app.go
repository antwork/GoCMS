package controllers

//后台首页
import "strconv"
import "github.com/robfig/revel"
import "admin/app/models"

type App struct {
	*revel.Controller
}

//首页
func (c App) Index(admin *models.Admin) revel.Result {
	title := "首页--GoCMS管理系统"

	if UserID, ok := c.Session["UserID"]; ok {
		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin_info := admin.GetById(UserID)
		if admin_info.Id <= 0 {
			return c.Redirect("/User/Login")
		}

		//控制器
		c.RenderArgs["Controller"] = c.Name
		//动作
		c.RenderArgs["action"] = c.Action
		//模型
		c.RenderArgs["Model"] = c.MethodName

		//导航菜单
		menu := new(models.Menu)
		c.RenderArgs["menus"] = menu.GetAdminMenu(0, false)

		//登陆用户信息
		c.RenderArgs["admin_info"] = admin_info

		//是否锁屏
		if c.Session["lock_screen"] == "" || c.Session["lock_screen"] == "0" {
			c.RenderArgs["lock_screen"] = "0"
		} else {
			c.RenderArgs["lock_screen"] = "1"
		}

	} else {
		return c.Redirect("/User/Login/")
	}

	c.Render(title)
	return c.RenderTemplate("App/Index.html")
}

func (c App) Main(admin *models.Admin) revel.Result {

	title := "首页--GoCMS管理系统"

	c.Render(title)

	return c.RenderTemplate("App/Main.html")
}
