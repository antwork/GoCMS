package controllers

//import "fmt"

import "github.com/robfig/revel"

//验证码
import "github.com/dchest/captcha"

type Captcha struct {
	*revel.Controller
}

//首页
func (c *Captcha) Index() revel.Result {
	captcha.Server(250, 62)
	CaptchaId := captcha.NewLen(6)
	captcha.WriteImage(c.Response.Out, CaptchaId, 250, 62)
	return nil
}
