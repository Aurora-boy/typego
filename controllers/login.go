package controllers

import (
	"typego/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "admin/login.html"
}

func (this *LoginController) DoLogin() {
	username := this.GetString("name")
	password := this.GetString("password")

	if username == "" {
		logs.Info("username is invialid")
	}

	orm := orm.NewOrm()

	user := models.User{}
	user.Username = username
	err := orm.Read(&user, "username")
	if err != nil {
		logs.Error("user not found!")
		return
	}
	logs.Info("find user: ", user.Username)
	if user.Password != password {
		logs.Info("password wrong!")
		return
	}
	this.Ctx.WriteString("登录成功")
}
