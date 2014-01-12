package models

//初始化模型

import "fmt"
import "github.com/robfig/config"
import _ "github.com/go-sql-driver/mysql"
import "github.com/lunny/xorm"
import "github.com/robfig/revel"
import "path/filepath"

var Engine *xorm.Engine

func init() {
	SetDB()
}

//设置数据库
func SetDB() {
	path, _ := filepath.Abs("")
	c, _ := config.ReadDefault(fmt.Sprintf("%s/admin/conf/databases.conf", path))

	driver, _ := c.String("database", "db.driver")
	dbname, _ := c.String("database", "db.dbname")
	user, _ := c.String("database", "db.user")
	password, _ := c.String("database", "db.password")
	host, _ := c.String("database", "db.host")
	//prefix, _ := c.String("database", "db.prefix")

	//数据库链接
	var err error
	Engine, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, host, dbname))
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//缓存方式是存放到内存中，缓存struct的记录数为1000条
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//Engine.SetDefaultCacher(cacher)

	//控制台打印SQL语句
	//Engine.ShowSQL = true

	//控制台打印调试信息
	//Engine.ShowDebug = true

	//控制台打印错误信息
	//Engine.ShowErr = true

	//控制台打印警告信息
	//Engine.ShowWarn = true
}
