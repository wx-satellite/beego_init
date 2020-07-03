package controllers

import (
	"byn/errors"
	"byn/requests"
	"byn/resources"
)

var userResources = map[int]resources.User{
	1: {"张三", 12},
	2: {"李四", 13},
	3: {"王五", 14},
}

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.Fail(errors.ErrorTkNotFound)
		return
	}
	user := userResources[id]
	if "" == user.Name {
		c.Fail(errors.ErrorTkNotFound)
		return
	}
	c.Success("获取成功", user)
	return
}

func (c *UserController) Post() {
	var form requests.User
	c.Validate(&form)
	max := 0
	for k := range userResources {
		if k > max {
			max = k
		}
	}
	userResources[max+1] = resources.User{Name: form.Name, Age: form.Age}
	c.Success("创建成功", nil)
	return
}

func (c *UserController) Home() {
	c.Data["json"] = struct {
		Message string
	}{Message: "Hello World"}
	c.ServeJSON()
}
