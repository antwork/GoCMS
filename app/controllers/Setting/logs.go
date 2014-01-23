package controllers

//日志管理
import "strconv"
import "admin/app/models"
import "github.com/robfig/revel"

type Logs struct {
	*revel.Controller
}

func (c Logs) Index(logs *models.Logs) revel.Result {

	title := "日志管理--GoCMS管理系统"

	var page string = c.Params.Get("page")

	where := make(map[string]string)

	var username string = c.Params.Get("username")
	if len(username) > 0 {
		where["username"] = username
	}
	var realname string = c.Params.Get("realname")
	if len(realname) > 0 {
		where["realname"] = realname
	}

	var start_time string = c.Params.Get("start_time")
	if len(start_time) > 0 {
		where["start_time"] = start_time
	}

	var end_time string = c.Params.Get("end_time")
	if len(end_time) > 0 {
		where["end_time"] = end_time
	}

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		logs_list, pages := logs.GetByAll(where, Page, 10)

		c.Render(title, logs_list, pages)

	} else {

		logs_list, pages := logs.GetByAll(where, 1, 10)

		c.Render(title, logs_list, pages)
	}

	return c.RenderTemplate("Setting/Logs/Index.html")
}
