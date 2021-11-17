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
	orm := orm.NewOrm()

	user := models.User{}

	// get information from form data
	user.Username = this.GetString("username")
	user.Password = this.GetString("password")
	user.Mail = this.GetString("mail")
	user.Created = time.Now()
	user.Logged = time.Now()

	// insert
	_, err := orm.Insert(&user)
	if err != nil {
		logs.Info("Insert error", err)
		return
	}
	this.Ctx.WriteString("注册成功")

}
