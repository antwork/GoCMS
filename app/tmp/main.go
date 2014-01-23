// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	_ "admin/app"
	controllers "admin/app/controllers"
	controllers0 "admin/app/controllers/Content"
	controllers1 "admin/app/controllers/Extend"
	controllers2 "admin/app/controllers/Module"
	controllers3 "admin/app/controllers/Panel"
	controllers4 "admin/app/controllers/Plugin"
	controllers5 "admin/app/controllers/Public"
	controllers6 "admin/app/controllers/Setting"
	controllers7 "admin/app/controllers/Style"
	controllers8 "admin/app/controllers/User"
	models "admin/app/models"
	tests "admin/tests"
	controllers9 "github.com/robfig/revel/modules/static/app/controllers"
	_ "github.com/robfig/revel/modules/testrunner/app"
	controllers10 "github.com/robfig/revel/modules/testrunner/app/controllers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					52: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Main",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					60: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Init)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*controllers0.Content)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Extend)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Module)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers3.Panel)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers4.Plugin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Ajax)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					15: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Pos",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetMessage",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ScreenLock",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ScreenUnlock",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Captcha)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetCaptchaId",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Public)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Map",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					17: []string{ 
						"map_html",
					},
				},
			},
			&revel.MethodType{
				Name: "CreateHtml",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					23: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Search",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					29: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Message",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					35: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Admin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					28: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					32: []string{ 
						"title",
						"admin_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					47: []string{ 
						"title",
						"role_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					179: []string{ 
						"title",
						"admin_info",
						"role_list",
					},
					181: []string{ 
						"title",
						"role_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Logs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "logs", Type: reflect.TypeOf((**models.Logs)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					47: []string{ 
						"title",
						"logs_list",
						"pages",
					},
					53: []string{ 
						"title",
						"logs_list",
						"pages",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Menu)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					18: []string{ 
						"title",
						"menus",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
						"title",
						"menus",
						"Id",
					},
					42: []string{ 
						"title",
						"menus",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					138: []string{ 
						"title",
						"menus",
						"menu_info",
					},
					144: []string{ 
						"title",
						"menus",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Role)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					26: []string{ 
						"title",
						"role_list",
						"pages",
					},
					30: []string{ 
						"title",
						"role_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Member",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					61: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					65: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					69: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					84: []string{ 
						"title",
						"tree",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					158: []string{ 
						"title",
						"role_info",
						"tree",
						"Id",
					},
					164: []string{ 
						"title",
						"tree",
					},
				},
			},
			&revel.MethodType{
				Name: "SetStatus",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Setting)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					14: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers7.Style)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers8.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					18: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					29: []string{ 
						"title",
						"CaptchaId",
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "EditInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					151: []string{ 
						"title",
						"admin_info",
					},
					153: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "EditPwd",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					214: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Left",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					306: []string{ 
						"title",
						"left_menu",
					},
					308: []string{ 
						"title",
					},
					326: []string{ 
						"title",
						"left_menu",
					},
					328: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers9.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers10.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
						"error",
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"admin/app/models.(*Admin).Validate": { 
			31: "a.Username",
			32: "a.Username",
			46: "a.Email",
			47: "a.Email",
			61: "a.Password",
			62: "a.Password",
		},
		"admin/app/models.(*Menu).Validate": { 
			21: "menu.Name",
			22: "menu.Name",
			23: "menu.Pid",
			24: "menu.Url",
			25: "menu.Order",
		},
		"admin/app/models.(*Password).ValidatePassword": { 
			67: "P.Password",
			68: "P.PasswordConfirm",
			70: "P.Password",
			71: "P.Password",
		},
	}
	revel.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
