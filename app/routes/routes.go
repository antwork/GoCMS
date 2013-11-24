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


type tCaptcha struct {}
var Captcha tCaptcha


func (_ tCaptcha) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.Index", args).Url
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
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Login", args).Url
}

func (_ tUser) LoginPost(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.LoginPost", args).Url
}

func (_ tUser) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Logout", args).Url
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


