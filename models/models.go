package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string
	Password string
	Mail     string
	Created  time.Time
	Logged   time.Time
}

func init() {
	driverName, _ := web.AppConfig.String("driverName")

	// connect to mysql
	orm.RegisterDriver(driverName, orm.DRMySQL)
	// mysql config
	user, _ := web.AppConfig.String("mysqlUser")
	pwd, _ := web.AppConfig.String("mysqlPwd")
	dbName, _ := web.AppConfig.String("dbName")
	host, _ := web.AppConfig.String("host")
	port, _ := web.AppConfig.String("port")

	// typego:typego@tcp(127.0.0.1:3306)/typego?charset=utf8
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		logs.Error("failed to connect to mysql database")
		return
	}
	logs.Info("connect to mysql sucessfully!")

	// register models
	orm.RegisterModel(new(User))
	// syncdb
	orm.RunSyncdb("default", false, true)
}
