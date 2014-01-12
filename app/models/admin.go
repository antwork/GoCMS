package models

//管理员表
import "admin/lib"
import "time"
import "regexp"
import "github.com/robfig/revel"

type Admin struct {
	Id            int64  `xorm:"pk autoincr"`
	Username      string `xorm:"unique index varchar(255)"`
	Password      string `xorm:"varchar:(32)"`
	Roleid        int64  `xorm:"index"`
	Role          *Role  `xorm:"- <- ->"`
	Lastloginip   string `xorm:"varchar(32)"`
	Lastlogintime string `xorm:"varchar(32)"`
	Email         string `xorm:"varchar(32)"`
	Realname      string `xorm:"varchar(32)"`
	Lang          string `xorm:"varchar(6)"`
}

type Password struct {
	Password        string
	PasswordConfirm string
}

func (a *Admin) Validate(v *revel.Validation) {
	v.Required(a.Username).Message("请输入用户名")
	valid := v.Match(a.Username, regexp.MustCompile("^\\w*$")).Message("只能使用字母、数字和下划线")
	if valid.Ok {
		if a.HasName() {
			err := &revel.ValidationError{
				Message: "该用户名已经注册过",
				Key:     "a.Username",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(a.Email).Message("请输入Email")
	valid = v.Email(a.Email).Message("无效的电子邮件")
	if valid.Ok {
		if a.HasEmail() {
			err := &revel.ValidationError{
				Message: "该邮件已经注册过",
				Key:     "a.Email",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(a.Password).Message("请输入密码")
	v.MinSize(a.Password, 3).Message("密码最少三位")
}

//验证密码
func (P *Password) ValidatePassword(v *revel.Validation) {
	v.Required(P.Password).Message("请输入密码")
	v.Required(P.PasswordConfirm).Message("请输入确认密码")

	v.MinSize(P.Password, 6).Message("密码最少六位")
	v.Required(P.Password == P.PasswordConfirm).Message("两次密码不相同!")
}

//获取管理员列表
func (a *Admin) GetByAll(RoleId int64) []*Admin {
	admin_list := []*Admin{}

	if RoleId > 0 {
		Engine.Where("roleid=?", RoleId).Find(&admin_list)
	} else {
		Engine.Find(&admin_list)
	}

	if len(admin_list) > 0 {
		role := new(Role)

		for i, v := range admin_list {
			admin_list[i].Role = role.GetById(v.Roleid)
		}
	}

	return admin_list
}

func (a *Admin) HasName() bool {

	admin := new(Admin)
	has, err := Engine.Where("email=?", a.Username).Get(admin)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if admin.Id > 0 {
		return true
	}
	return false
}

func (a *Admin) HasEmail() bool {

	admin := new(Admin)
	has, err := Engine.Where("email=?", a.Email).Get(admin)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if admin.Id > 0 {
		return true
	}
	return false
}

//根据Id获取管理员信息
func (a *Admin) GetById(Id int64) *Admin {
	admin := new(Admin)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := Engine.Id(Id).Get(admin)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return admin
}

//根据用户名获取管理员信息
func (a *Admin) GetByName(name string) *Admin {
	admin := new(Admin)
	has, err := Engine.Where("username=?", name).Get(admin)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return admin
}

//添加管理员
func (a *Admin) Save() bool {

	admin := new(Admin)
	admin.Username = a.Username
	admin.Password = lib.Md5(a.Password)
	admin.Roleid = a.Roleid
	admin.Lastloginip = lib.GetClientIP()
	admin.Email = a.Email
	admin.Realname = a.Realname
	admin.Lang = a.Lang
	admin.Lastlogintime = time.Now().Format("2006-01-02 15:04:04")

	has, err := Engine.Insert(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑管理员
func (a *Admin) Edit(Id int64) bool {

	admin := new(Admin)

	if len(a.Username) > 0 {
		admin.Username = a.Username
	}

	if len(a.Password) > 0 {
		admin.Password = lib.Md5(a.Password)
	}

	if a.Roleid > 0 {
		admin.Roleid = a.Roleid
	}

	if len(a.Email) > 0 {
		admin.Email = a.Email
	}

	if len(a.Realname) > 0 {
		admin.Realname = a.Realname
	}

	if len(a.Lang) > 0 {
		admin.Lang = a.Lang
	}

	admin.Lastloginip = lib.GetClientIP()
	admin.Lastlogintime = time.Now().Format("2006-01-02 15:04:04")

	has, err := Engine.Id(Id).Update(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//删除管理员
func (a *Admin) DelByID(Id int64) bool {

	admin := new(Admin)

	has, err := Engine.Id(Id).Delete(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
