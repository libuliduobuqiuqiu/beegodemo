package main

import (
	_ "beegodemo/api"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
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

type Device struct {
	Name    string `json:"name"`
	Address string `json:"addess"`
}

func GetDevice(ctx *context.Context) {
	d := Device{
		Name:    "server1",
		Address: "localhost",
	}
	ctx.JSONResp(d)
}

func main() {
	web.CtrlGet("/index", (*IndexController).Index)
	web.Get("/device", GetDevice)
	web.Run(":8090")
}
