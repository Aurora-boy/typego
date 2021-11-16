package main

import (
	_ "typego/models"
	_ "typego/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
