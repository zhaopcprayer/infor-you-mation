package main

import (
	"flag"
	"github.com/astaxie/beego"
	_ "github.com/yanyiwu/infor-you-mation/routers"
)

func main() {
	flag.Parse()
	//beego.SessionOn = true
	//beego.SetStaticPath("/publish", "statics")
	beego.Run()
}
