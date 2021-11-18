package controllers

import (
	"time"
	"typego/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "admin/register.html"
}

func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	mail := this.GetString("mail")

	orm := orm.NewOrm()

	user := models.User{}

	// get information from form data
	user.Username = username
	if orm.Read(&user, "username") != nil {
		// user didn't in users table
		user.Password = password
		user.Mail = mail
		user.Created = time.Now()
		user.Logged = time.Now()

		// insert
		_, err := orm.Insert(&user)
		if err != nil {
			logs.Info("Insert error", err)
			return
		}
		// this.Ctx.WriteString("注册成功")
	} else {
		logs.Info("user already exists")
	}
	this.Redirect("/login", 302)

}
