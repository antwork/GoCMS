package models

//日志管理
import "time"
import "admin/utils"
import "html/template"
import "github.com/robfig/revel"

type Logs struct {
	Id         int64  `xorm:"pk autoincr"`
	Uid        int64  `xorm:"unique"`
	Admin      *Admin `xorm:"- <- ->"`
	Module     string `xorm:"varchar:(50)"`
	Url        string `xorm:"varchar(100)"`
	Action     string `xorm:"varchar(100)"`
	Ip         string `xorm:"varchar(15)"`
	Desc       string `xorm:"text"`
	Createtime string `xorm:"DateTime"`
}

func (L *Logs) Validate(v *revel.Validation) {

}

//添加日志记录
func (L *Logs) Save(Admin_Info *Admin, c *revel.Controller, Desc string) bool {
	logs := new(Logs)

	logs.Uid = Admin_Info.Id
	logs.Module = c.Name
	logs.Url = c.Action
	logs.Action = c.MethodName
	logs.Ip = utils.GetClientIP()
	logs.Desc = Desc
	logs.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := Engine.Insert(logs)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//获取日志列表
func (L *Logs) GetByAll(where map[string]string, Page int64, Perpage int64) ([]*Logs, template.HTML) {

	logs_list := []*Logs{}

	//查询总数
	logs := new(Logs)
	Total, err := Engine.Count(logs)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Logs/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	Engine.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&logs_list)

	if len(logs_list) > 0 {
		admin := new(Admin)

		for i, v := range logs_list {
			logs_list[i].Admin = admin.GetById(v.Uid)
		}
	}

	return logs_list, pages
}
