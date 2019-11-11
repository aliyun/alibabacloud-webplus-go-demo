package main

import (
	_ "alibaba/com/webplusdemo/routers"
	"alibaba/com/webplusdemo/services"
	"github.com/astaxie/beego"
)


func main() {
	services.LoadLocales()
	beego.Run()
}
