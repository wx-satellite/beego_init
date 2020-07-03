package controllers

import (
	"byn/errors"
	"byn/resources"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Success(message string, data interface{}) {
	c.Data["json"] = resources.GetSuccessResponse(message, data)
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) Fail(err error) {
	c.Data["json"] = resources.GetFailResponse(err)
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) Validate(request interface{}) {
	// 绑定数据
	err := c.ParseForm(request)
	if err != nil {
		c.Fail(err)
		return
	}
	// 验证
	v := validation.Validation{}
	isValid, err := v.Valid(request)
	if err != nil {
		c.Fail(err)
		return
	}
	// 出现错误，只取第一个错误
	if !isValid {
		for _, e := range v.Errors {
			c.Fail(&errors.BynError{Code: errors.ValidateErrorCode, Message: e.Message})
			return
		}
	}
}
