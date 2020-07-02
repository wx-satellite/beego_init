package sysinit

import (
	_ "byn/models" // 可能模型也有一些初始化操作，例如注册模型等等
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MaxOpen = 200
	MaxIdle = 150
)

func dbinit(aliases ...string) {
	runMode := beego.AppConfig.String("runmode")
	isDev := "dev" == runMode
	if len(aliases) >= 2 {
		for _, alias := range aliases {
			registerMysqlDatabase(alias)
			if "w" == alias {
				_ = orm.RunSyncdb("default", false, isDev)
			}
		}
	} else {
		registerMysqlDatabase("w")
		_ = orm.RunSyncdb("default", false, isDev)
	}

	// 开启调试
	orm.Debug = isDev
}

func registerMysqlDatabase(alias string) {
	name := "default"
	if "w" != alias {
		name = alias
	}

	host := beego.AppConfig.String(fmt.Sprintf("mysql::%v_host", alias))
	port := beego.AppConfig.String(fmt.Sprintf("mysql::%v_port", alias))
	username := beego.AppConfig.String(fmt.Sprintf("mysql::%v_username", alias))
	password := beego.AppConfig.String(fmt.Sprintf("mysql::%v_password", alias))
	dbName := beego.AppConfig.String(fmt.Sprintf("mysql::%v_database", alias))

	err := orm.RegisterDataBase(name, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbName), MaxIdle, MaxOpen)
	if err != nil {
		logs.Error("数据库注册失败", err)
	}
}
