package main

import (
	"bee-api/handlers"
	_ "bee-api/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// init log
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)

	// get database configuration from environment variables
	dbUser, _ := beego.AppConfig.String("dbUser")
	dbPwd, _ := beego.AppConfig.String("dbPwd")
	dbName, _ := beego.AppConfig.String("dbName")
	dbString := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=utf8"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbString)
	orm.RunCommand()

}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.ErrorController(&handlers.ErrorController{})
	beego.Run("localhost")
}
