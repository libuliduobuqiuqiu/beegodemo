package main

import (
	_ "beegodemo/api"
	"github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	web.Controller
}

type Resp struct {
	Code int    `json:"resp"`
	Info string `json:"info"`
}

func (i *IndexController) Index() {
	resp := Resp{
		Code: 200,
		Info: "hello,world",
	}
	i.Data["json"] = resp
	i.ServeJSON()
}

func main() {
	web.BConfig.CopyRequestBody = true
	web.CtrlGet("/index", (*IndexController).Index)
	web.Run(":8090")
}
