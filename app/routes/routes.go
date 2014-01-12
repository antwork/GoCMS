// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Main(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Main", args).Url
}


type tInit struct {}
var Init tInit



type tContent struct {}
var Content tContent


func (_ tContent) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Content.Index", args).Url
}


type tExtend struct {}
var Extend tExtend


func (_ tExtend) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Extend.Index", args).Url
}


type tLogs struct {}
var Logs tLogs


func (_ tLogs) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Logs.Index", args).Url
}


type tModule struct {}
var Module tModule


func (_ tModule) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Module.Index", args).Url
}


type tPanel struct {}
var Panel tPanel


func (_ tPanel) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Panel.Index", args).Url
}


type tPlugin struct {}
var Plugin tPlugin


func (_ tPlugin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Plugin.Index", args).Url
}


type tAjax struct {}
var Ajax tAjax


func (_ tAjax) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.Login", args).Url
}

func (_ tAjax) Pos(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Ajax.Pos", args).Url
}

func (_ tAjax) GetMessage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.GetMessage", args).Url
}

func (_ tAjax) ScreenLock(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.ScreenLock", args).Url
}

func (_ tAjax) ScreenUnlock(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Ajax.ScreenUnlock", args).Url
}


type tCaptcha struct {}
var Captcha tCaptcha


func (_ tCaptcha) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.Index", args).Url
}


type tPublic struct {}
var Public tPublic


func (_ tPublic) Map(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Public.Map", args).Url
}

func (_ tPublic) CreateHtml(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.CreateHtml", args).Url
}

func (_ tPublic) Search(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.Search", args).Url
}

func (_ tPublic) Message(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.Message", args).Url
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Index", args).Url
}

func (_ tAdmin) Add(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Add", args).Url
}

func (_ tAdmin) Edit(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Edit", args).Url
}

func (_ tAdmin) Delete(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Delete", args).Url
}


type tMenu struct {}
var Menu tMenu


func (_ tMenu) Index(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Index", args).Url
}

func (_ tMenu) Add(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Add", args).Url
}

func (_ tMenu) Edit(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Edit", args).Url
}

func (_ tMenu) Delete(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Delete", args).Url
}


type tRole struct {}
var Role tRole


func (_ tRole) Index(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Index", args).Url
}

func (_ tRole) Member(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Member", args).Url
}

func (_ tRole) Add(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Add", args).Url
}

func (_ tRole) Edit(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Edit", args).Url
}

func (_ tRole) SetStatus(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.SetStatus", args).Url
}

func (_ tRole) Delete(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Delete", args).Url
}


type tSetting struct {}
var Setting tSetting


func (_ tSetting) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Setting.Index", args).Url
}


type tStyle struct {}
var Style tStyle


func (_ tStyle) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Style.Index", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Index", args).Url
}

func (_ tUser) Login(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.Login", args).Url
}

func (_ tUser) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Logout", args).Url
}

func (_ tUser) EditInfo(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditInfo", args).Url
}

func (_ tUser) EditPwd(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditPwd", args).Url
}

func (_ tUser) Left(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("User.Left", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


