package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel()
}

func TablePrefix() string {
	prefix := beego.AppConfig.String("table_prefix")
	return prefix
}

func GetUserTableName() string {
	return TablePrefix() + "users"
}
