package controllers

//初始化入口文件
import "runtime"
import "strconv"
import "github.com/robfig/revel"
import "admin/app/models"

type Init struct {
	*revel.Controller
}

func init() {

	revel.OnAppStart(func() {
		revel.WARN.Println("开始执行")

		//检测是否登陆
		revel.InterceptMethod(CheckLogin, revel.BEFORE)

		//多核运行
		np := runtime.NumCPU()
		if np >= 2 {
			runtime.GOMAXPROCS(np - 1)
		}
	})
}

//检测登陆
func CheckLogin(c *Init) revel.Result {
	if c.Name == "User" && c.MethodName == "Login" {
		return c.Render()
	} else {
		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
				return c.Redirect("/User/Login/")
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)
			if admin_info.Id <= 0 {
				return c.Redirect("/User/Login/")
			}
		} else {
			return c.Redirect("/User/Login/")
		}
	}

	return nil
}
