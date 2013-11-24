package models

//菜单管理
import "github.com/robfig/revel"

type Menu struct {
	Id         int64  `xorm:"pk"`
	Name       string `xorm:"unique varchar(255)"`
	Pid        int64
	Controller string `xorm:"varchar(60)"`
	Data       string `xorm:"varchar(60)"`
	Listorder  int16  `xorm:"-"`
	Display    int8   `xorm:"default(1)"`
}

func init() {
}

func (m *Menu) Validate(v *revel.Validation) {
	v.Required(m.Name).Message("请输入菜单名称")
	v.Required(m.Pid).Message("请选择父菜单")
	v.Required(m.Controller).Message("请输入菜单地址")

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
	menu.Controller = m.Controller
	menu.Data = m.Data
	menu.Listorder = m.Listorder
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

	if len(m.Controller) > 0 {
		menu.Controller = m.Controller
	}

	if len(m.Data) > 0 {
		menu.Data = m.Data
	}

	if m.Listorder > 0 {
		menu.Listorder = m.Listorder
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
//Pid 父菜单ID
//是否包括他自己
func (m *Menu) GetMenuAll(Pid int64, with_self bool) []*Menu {
	menus := []*Menu{}
	Engine.Where("Pid=?", Pid).Find(&menus)
	return menus
}
