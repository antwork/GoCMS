package controllers

//AJAX操作
import "strconv"
import "admin/utils"
import "github.com/robfig/revel"
import "admin/app/models"

type Ajax struct {
	*revel.Controller
}

//登陆
func (c *Ajax) Login() revel.Result {
	return c.Render()
}

//当前位置
func (c *Ajax) Pos(menu *models.Menu) revel.Result {
	var id string = c.Params.Get("id")

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	menu_str := menu.GetPos(Id)
	return c.RenderText(menu_str)
}

//检查消息
func (c *Ajax) GetMessage() revel.Result {
	data := make(map[string]string)

	data["status"] = "0"
	data["message"] = "请填写用户名!"
	return c.RenderJson(data)
}

//锁屏
func (c *Ajax) ScreenLock() revel.Result {
	data := make(map[string]string)

	c.Session["lock_screen"] = "1"

	data["status"] = "1"
	data["message"] = "锁屏!"
	return c.RenderJson(data)
}

//解锁
func (c *Ajax) ScreenUnlock(admin *models.Admin) revel.Result {
	var lock_password string = c.Params.Get("lock_password")

	if lock_password == "" || len(lock_password) <= 0 {
		return c.RenderText("2")
	}

	if UserID, ok := c.Session["UserID"]; ok {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin_info := admin.GetById(UserID)

		if admin_info.Password != utils.Md5(lock_password) {
			return c.RenderText("3")
		} else {
			c.Session["lock_screen"] = "0"
			return c.RenderText("1")
		}
	} else {
		return c.RenderText("4")
	}
}
