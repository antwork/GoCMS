package controllers

//后台登陆
import "fmt"
import "admin/lib"
import "github.com/robfig/revel"
import "admin/app/models"
import "admin/app/routes"

type User struct {
	*revel.Controller
}

//首页
func (c *User) Index() revel.Result {
	return c.Render()
}

//登陆
func (c *User) Login() revel.Result {
	return c.Render()
}

//后台登陆
func (c *User) LoginPost() revel.Result {

	var username string = c.Params.Get("username")
	var password string = c.Params.Get("password")
	var verify string = c.Params.Get("verify")

	data := make(map[string]string)

	if len(username) <= 0 {
		data["status"] = "0"
		data["url"] = "/"
		data["info"] = "请填写用户名!"
		return c.RenderJson(data)
	}

	if len(password) <= 0 {
		data["status"] = "0"
		data["url"] = "/"
		data["info"] = "请填写密码!"
		return c.RenderJson(data)
	}

	if len(verify) <= 0 {
		data["status"] = "0"
		data["url"] = "/"
		data["info"] = "请填写验证码!"
		return c.RenderJson(data)
	}

	admin := new(models.Admin)

	admin = admin.GetByName(username)

	if admin.Id <= 0 {
		data["status"] = "0"
		data["url"] = "/"
		data["info"] = "用户名错误!"
	} else if username == admin.Username && lib.EncryptPassword(password) == admin.Password {
		c.Session["UserID"] = fmt.Sprintf("%d", admin.Id)

		data["status"] = "1"
		data["url"] = "/"
		data["info"] = "登陆成功!"
	} else {
		data["status"] = "0"
		data["url"] = "/"
		data["info"] = "密码错误!"
	}

	return c.RenderJson(data)
}

//退出登陆
func (c *User) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.User.Login())
}
