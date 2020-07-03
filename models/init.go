package models

import (
	"github.com/astaxie/beego"
)

func init() {
	//orm.RegisterModel(new(User))
}

func TablePrefix() string {
	prefix := beego.AppConfig.String("table_prefix")
	return prefix
}

func GetUserTableName() string {
	return TablePrefix() + "users"
}
