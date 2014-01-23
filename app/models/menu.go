package models

//菜单管理
import "strconv"
import "strings"
import "admin/utils"
import "html/template"
import "github.com/robfig/revel"

type Menu struct {
	Id      int64  `xorm:"pk"`
	Name    string `xorm:"unique varchar(255)"`
	Pid     int64  `xorm:"int(11)"`
	Url     string `xorm:"char(100)"`
	Data    string `xorm:"varchar(60)"`
	Order   int64  `xorm:"int(11)"`
	Display int64  `xorm:"default 1"`
}

func (menu *Menu) Validate(v *revel.Validation) {
	v.Required(menu.Name).Message("请输入菜单名称!")
	v.MaxSize(menu.Name, 105).Message("最多35个字")
	v.Required(menu.Pid).Message("请选择父菜单!")
	v.Required(menu.Url).Message("请输入菜单地址!")
	v.Required(menu.Order).Message("请输入排序!")
}

//根据Id获取菜单信息
func (m *Menu) GetById(Id int64) *Menu {

	menu := new(Menu)
	has, err := Engine.Id(Id).Get(menu)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return menu
}

//添加菜单
func (m *Menu) Save() bool {

	menu := new(Menu)
	menu.Name = m.Name
	menu.Pid = m.Pid
	menu.Url = m.Url
	menu.Data = m.Data
	menu.Order = m.Order
	menu.Display = m.Display

	has, err := Engine.Insert(menu)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑菜单
func (m *Menu) Edit(Id int64) bool {
	menu := new(Menu)

	if len(m.Name) > 0 {
		menu.Name = m.Name
	}

	if m.Pid > 0 {
		menu.Pid = m.Pid
	}

	if len(m.Url) > 0 {
		menu.Url = m.Url
	}

	if len(m.Data) > 0 {
		menu.Data = m.Data
	}

	if m.Order > 0 {
		menu.Order = m.Order
	}

	if m.Display > 0 {
		menu.Display = m.Display
	}

	has, err := Engine.Id(Id).Update(menu)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//删除菜单
func (m *Menu) DelByID(Id int64) bool {
	menu := new(Menu)

	has, err := Engine.Id(Id).Delete(menu)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//按父ID查找菜单子项
func (m *Menu) GetAdminMenu(Pid int64, Admin_Info *Admin) []*Menu {

	//初始化菜单
	menus := []*Menu{}

	if Admin_Info.Id != 1 {
		err := Engine.Where("Pid=? AND Display=? and id in ("+Admin_Info.Role.Data+")", Pid, 1).Find(&menus)

		if err != nil {
			revel.WARN.Printf("错误: %v", err)
		}
	} else {
		err := Engine.Where("Pid=? AND Display=?", Pid, 1).Find(&menus)

		if err != nil {
			revel.WARN.Printf("错误: %v", err)
		}
	}

	return menus
}

//获取所有菜单
func (m *Menu) GetMenuAll() map[int64][]*Menu {

	menus := make([]*Menu, 0)
	Engine.Asc("order").Find(&menus)

	//初始化菜单Map
	menu_list := make(map[int64][]*Menu)

	for _, menu := range menus {
		if _, ok := menu_list[menu.Pid]; !ok {
			menu_list[menu.Pid] = make([]*Menu, 0)
		}
		menu_list[menu.Pid] = append(menu_list[menu.Pid], menu)
	}

	return menu_list
}

//获取左侧导航菜单
func (m *Menu) GetLeftMenuHtml(Pid int64, Admin_Info *Admin) template.HTML {

	menus := make([]*Menu, 0)

	if Admin_Info.Id != 1 {
		err := Engine.Where("id in (" + Admin_Info.Role.Data + ")").Asc("order").Find(&menus)

		if err != nil {
			revel.WARN.Printf("错误: %v", err)
		}
	} else {
		err := Engine.Asc("order").Find(&menus)

		if err != nil {
			revel.WARN.Printf("错误: %v", err)
		}
	}

	//初始化菜单Map
	menu_list := make(map[int64][]*Menu)

	for _, menu := range menus {
		if _, ok := menu_list[menu.Pid]; !ok {
			menu_list[menu.Pid] = make([]*Menu, 0)
		}
		menu_list[menu.Pid] = append(menu_list[menu.Pid], menu)
	}

	Html := ""
	for _, menu_second := range menu_list[Pid] {
		Html += "<h3 class=\"f14\"><span class=\"switchs cu on\" title=\"展开与收缩\"></span>" + menu_second.Name + "</h3>"

		Html += "<ul>"
		for _, menu_last := range menu_list[menu_second.Id] {
			Html += "<li id=\"_MP" + strconv.FormatInt(menu_last.Id, 10) + "\" class=\"sub_menu\">"
			Html += "<a href=\"javascript:_MP(" + strconv.FormatInt(menu_last.Id, 10) + ",'" + menu_last.Url + "');\" hidefocus=\"true\" style=\"outline:none;\">" + menu_last.Name + "</a>"
			Html += "</li>"
		}
		Html += "</ul>"
	}

	return template.HTML(Html)
}

//返回菜单Option的HTML
func (m *Menu) GetMenuOptionHtml(Id int64) template.HTML {
	menus := make([]*Menu, 0)
	Engine.Asc("order").Find(&menus)

	//初始化菜单Map
	menu_list := make(map[int64][]*Menu)

	for _, menu := range menus {
		if _, ok := menu_list[menu.Pid]; !ok {
			menu_list[menu.Pid] = make([]*Menu, 0)
		}
		menu_list[menu.Pid] = append(menu_list[menu.Pid], menu)
	}

	Html := ""
	for _, menu := range menu_list[0] {

		if menu.Id == Id {
			Html += "<option value=" + strconv.FormatInt(menu.Id, 10) + " selected ><b>" + menu.Name + "</b></option>"
		} else {
			Html += "<option value=" + strconv.FormatInt(menu.Id, 10) + "><b>" + menu.Name + "</b></option>"
		}

		for _, menu_second := range menu_list[menu.Id] {
			if menu_second.Id == Id {
				Html += "<option value=" + strconv.FormatInt(menu_second.Id, 10) + " selected >&#12288;&#8866;" + menu_second.Name + "</option>"
			} else {
				Html += "<option value=" + strconv.FormatInt(menu_second.Id, 10) + ">&#12288;&#8866;" + menu_second.Name + "</option>"
			}

			for _, menu_last := range menu_list[menu_second.Id] {
				if menu_last.Id == Id {
					Html += "<option value=" + strconv.FormatInt(menu_last.Id, 10) + " selected >&#12288;&#12288;&#8866;" + menu_last.Name + "</option>"
				} else {
					Html += "<option value=" + strconv.FormatInt(menu_last.Id, 10) + ">&#12288;&#12288;&#8866;" + menu_last.Name + "</option>"
				}

			}
		}
	}

	return template.HTML(Html)
}

//返回菜单树
func (m *Menu) GetMenuTree(role string) template.HTML {
	menus := make([]*Menu, 0)
	Engine.Asc("order").Find(&menus)

	//初始化菜单Map
	menu_list := make(map[int64][]*Menu)

	for _, menu := range menus {
		if _, ok := menu_list[menu.Pid]; !ok {
			menu_list[menu.Pid] = make([]*Menu, 0)
		}
		menu_list[menu.Pid] = append(menu_list[menu.Pid], menu)
	}

	//解析权限
	arr_role := strings.Split(role, ",")

	Html := "<SCRIPT type=\"text/javascript\">var zNodes =["

	for _, menu := range menu_list[0] {

		if utils.In_Array(strconv.FormatInt(menu.Id, 10), arr_role) {
			Html += "{ id:" + strconv.FormatInt(menu.Id, 10) + ", pId:" + strconv.FormatInt(menu.Pid, 10) + ", name:'" + menu.Name + "', open:true, checked:true},"
		} else {
			Html += "{ id:" + strconv.FormatInt(menu.Id, 10) + ", pId:" + strconv.FormatInt(menu.Pid, 10) + ", name:'" + menu.Name + "', open:true},"
		}

		for _, menu_second := range menu_list[menu.Id] {
			if utils.In_Array(strconv.FormatInt(menu_second.Id, 10), arr_role) {
				Html += "{ id:" + strconv.FormatInt(menu_second.Id, 10) + ", pId:" + strconv.FormatInt(menu_second.Pid, 10) + ", name:'" + menu_second.Name + "', open:true, checked:true},"
			} else {
				Html += "{ id:" + strconv.FormatInt(menu_second.Id, 10) + ", pId:" + strconv.FormatInt(menu_second.Pid, 10) + ", name:'" + menu_second.Name + "', open:true},"
			}

			for _, menu_last := range menu_list[menu_second.Id] {
				if utils.In_Array(strconv.FormatInt(menu_last.Id, 10), arr_role) {
					Html += "{ id:" + strconv.FormatInt(menu_last.Id, 10) + ", pId:" + strconv.FormatInt(menu_last.Pid, 10) + ", name:'" + menu_last.Name + "', checked:true},"
				} else {
					Html += "{ id:" + strconv.FormatInt(menu_last.Id, 10) + ", pId:" + strconv.FormatInt(menu_last.Pid, 10) + ", name:'" + menu_last.Name + "'},"
				}

			}
		}
	}

	Html += "];</SCRIPT>"

	return template.HTML(Html)
}

//返回后台地图
func (m *Menu) GetMenuMap() template.HTML {
	menus := make([]*Menu, 0)
	Engine.Asc("order").Find(&menus)

	//初始化菜单Map
	menu_list := make(map[int64][]*Menu)

	for _, menu := range menus {
		if _, ok := menu_list[menu.Pid]; !ok {
			menu_list[menu.Pid] = make([]*Menu, 0)
		}
		menu_list[menu.Pid] = append(menu_list[menu.Pid], menu)
	}

	Html := ""
	n := 1
	for _, menu := range menu_list[0] {
		if n == 1 {
			Html += "<div class=\"map-menu lf\">"
		}

		Html += "<ul>"
		Html += "<li class=\"title\">" + menu.Name + "</li>"

		for _, menu_second := range menu_list[menu.Id] {
			Html += "<li class=\"title2\">" + menu_second.Name + "</li>"

			for _, menu_last := range menu_list[menu_second.Id] {
				Html += "<li><a href=\"javascript:Go(" + strconv.FormatInt(menu_last.Id, 10) + ",'" + menu_last.Url + "')\">" + menu_last.Name + "</a></li>"
			}
		}

		Html += "</ul>"

		if n%2 == 0 {
			Html += "</div><div class=\"map-menu lf\">"
		}
		n++
	}

	return template.HTML(Html)
}

//获取所有菜单
//返回HTML
func (m *Menu) GetMenuHtml() template.HTML {
	menus := make([]*Menu, 0)
	Engine.Asc("order").Find(&menus)

	//初始化菜单Map
	menu_list := make(map[int64][]*Menu)

	for _, menu := range menus {
		if _, ok := menu_list[menu.Pid]; !ok {
			menu_list[menu.Pid] = make([]*Menu, 0)
		}
		menu_list[menu.Pid] = append(menu_list[menu.Pid], menu)
	}

	Html := ""
	for _, menu := range menu_list[0] {

		Html += "<tr>"
		Html += "<td align=\"left\"><b>" + menu.Name + "</b></td>"
		Html += "<td align=\"center\">" + strconv.FormatInt(menu.Order, 10) + "</td>"
		Html += "<td align=\"center\">" + menu.Url + "</td>"
		if menu.Display == 1 {
			Html += "<td align=\"center\">菜单显示</td>"
		} else {
			Html += "<td align=\"center\">菜单隐藏</td>"
		}

		Html += "<td align=\"center\">"
		Html += "<a href=\"/Menu/add/" + strconv.FormatInt(menu.Id, 10) + "/\">添加子菜单</a> |"
		Html += "<a href=\"/Menu/edit/" + strconv.FormatInt(menu.Id, 10) + "/\">修改</a> |"
		Html += "<a onclick=\"delete_menu(" + strconv.FormatInt(menu.Id, 10) + ")\" href=\"javascript:;\">删除</a>"
		Html += "</td>"

		Html += "</tr>"

		for _, menu_second := range menu_list[menu.Id] {
			Html += "<tr>"
			Html += "<td align=\"left\">&nbsp;&nbsp;&#12288;&#8866;&nbsp;&nbsp;" + menu_second.Name + "</td>"
			Html += "<td align=\"center\">" + strconv.FormatInt(menu_second.Order, 10) + "</td>"
			Html += "<td align=\"center\">" + menu_second.Url + "</td>"
			if menu_second.Display == 1 {
				Html += "<td align=\"center\">菜单显示</td>"
			} else {
				Html += "<td align=\"center\">菜单隐藏</td>"
			}

			Html += "<td align=\"center\">"
			Html += "<a href=\"/Menu/add/" + strconv.FormatInt(menu_second.Id, 10) + "/\">添加子菜单</a> |"
			Html += "<a href=\"/Menu/edit/" + strconv.FormatInt(menu_second.Id, 10) + "/\">修改</a> |"
			Html += "<a onclick=\"delete_menu(" + strconv.FormatInt(menu_second.Id, 10) + ")\" href=\"javascript:;\">删除</a>"
			Html += "</td>"

			Html += "</tr>"

			for _, menu_last := range menu_list[menu_second.Id] {
				Html += "<tr>"
				Html += "<td align=\"left\">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&#12288;&#8866;&nbsp;&nbsp;" + menu_last.Name + "</td>"
				Html += "<td align=\"center\">" + strconv.FormatInt(menu_last.Order, 10) + "</td>"
				Html += "<td align=\"center\">" + menu_last.Url + "</td>"
				if menu_last.Display == 1 {
					Html += "<td align=\"center\">菜单显示</td>"
				} else {
					Html += "<td align=\"center\">菜单隐藏</td>"
				}

				Html += "<td align=\"center\">"
				//Html += "<a href=\"/Menu/add/" + strconv.FormatInt(menu_last.Id, 10) + "/\">添加子菜单</a> |"
				Html += "<a href=\"/Menu/edit/" + strconv.FormatInt(menu_last.Id, 10) + "/\">修改</a> |"
				Html += "<a onclick=\"delete_menu(" + strconv.FormatInt(menu_last.Id, 10) + ")\" href=\"javascript:;\">删除</a>"
				Html += "</td>"

				Html += "</tr>"
			}
		}
	}

	return template.HTML(Html)
}

//当前位置
func (m *Menu) GetPos(Id int64) string {
	menu := new(Menu)
	has, err := Engine.Id(Id).Get(menu)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	var str string = ""

	if menu.Pid > 0 {
		str += m.GetPos(menu.Pid)
	}

	return str + menu.Name + " > "
}
