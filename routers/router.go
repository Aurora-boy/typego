package routers

import (
	"typego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// admin
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:DoLogin")
	beego.Router("/register", &controllers.RegisterController{}, "get:Get;post:Post")
}
