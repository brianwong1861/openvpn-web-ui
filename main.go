package main

import (
	"github.com/brianwong1861/openvpn-web-ui/lib"
	_ "github.com/brianwong1861/openvpn-web-ui/routers"
	"github.com/astaxie/beego"
)

func main() {
	lib.AddFuncMaps()
	beego.Run()
}
